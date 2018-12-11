package main // inport  "try-reuseport"

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gin "github.com/gin-gonic/gin"
	"golang.org/x/sys/unix"
)

var msg = "Hello Gopher"

func main() {
	h := gin.Default()
	h.GET("/", handler)

	lc := net.ListenConfig{
		Control: listenCtrl,
	}
	l, err := lc.Listen(context.Background(), "tcp4", ":8080")
	if err != nil {
		panic(err)
	}

	svc := http.Server{
		Addr:    ":8080",
		Handler: h,
	}
	go func() {
		if err = svc.Serve(l); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Fatalln("Server closed with error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := svc.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		log.Println("Failed to gracefully shutdown:", err)
	}
	log.Println("Server shutdown")
}

func handler(c *gin.Context) {
	time.Sleep(3 * time.Second)
	fmt.Println("This server")
	c.JSON(200, gin.H{
		"message": msg,
	})
}

func listenCtrl(network string, address string, c syscall.RawConn) error {
	var operr error
	var fn = func(s uintptr) {
		operr = unix.SetsockoptInt(int(s), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
	}
	if err := c.Control(fn); err != nil {
		return err
	}
	if operr != nil {
		return operr
	}
	return nil
}
