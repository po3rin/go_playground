package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func main() {
	srv := &http.Server{
		Addr: ":8080",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	errC := make(chan error)
	defer close(errC)

	errChanNum := 3
	go func() {
		errC <- run(ctx)
	}()
	go func() {
		errC <- run(ctx)
	}()
	go func() {
		errC <- srv.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("start!")

	// エラーかシグナルを受けたら 全ての gorutine をちゃんと止めて main を終了したい
	var once sync.Once
	var errChanCount int
	for {
		// シグナルかエラーを待つ
		select {
		case <-quit:
			fmt.Println("received signal")
		case err := <-errC:
			errChanCount++
			if err != nil {
				log.Println(err)
			}
		}

		// サーバーのシャットダウンとコンテキストキャンセル
		// 何回コールしても大丈夫だけど一応 sync.Once で
		once.Do(func() {
			fmt.Println("cancel & shutdown")
			cancel()
			if err := srv.Shutdown(ctx); err != nil {
				log.Println(err)
			}
		})

		// gorutineの終了をerrCの帰ってきた数で知るのはアンチパターン？
		// コンテキストキャンセルしたらgorutineで起動しているプロセスが
		// 確実に終了することが前提の実装になっている
		if errChanCount == errChanNum {
			break
		}
	}

	fmt.Println("done")
}
