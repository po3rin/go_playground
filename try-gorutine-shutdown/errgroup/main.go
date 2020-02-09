package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func httpServe(ctx context.Context) error {
	srv := &http.Server{Addr: ":8080"}

	// errorを返したらコンテキストを閉じる
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return srv.ListenAndServe()
	})

	// コンテキストキャンセルを受けたら
	<-ctx.Done()
	sCtx, sCancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer sCancel()
	if err := srv.Shutdown(sCtx); err != nil {
		return err
	}

	return eg.Wait()
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})
	eg.Go(func() error {
		return run(ctx)
	})
	eg.Go(func() error {
		return httpServe(ctx)
	})

	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("start!")

	// シグナルかコンテキストキャンセルを受ける
	select {
	case <-quit:
		cancel()
	case <-ctx.Done():
	}

	// コンテキストキャンセルしたら全てのgorutineが確実に終了することが前提の実装になっている
	if err := eg.Wait(); err != nil {
		log.Println(err)
	}

	fmt.Println("done")
}
