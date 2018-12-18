package main

import (
	"fmt"
	"go-playground/detect-nill-res/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", handler.Hello)
	fmt.Println("run API at 8080 port")
	r.Run(":8888")
}
