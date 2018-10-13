package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

//TokenKey for use token.
type TokenKey string

var key TokenKey

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8822", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 2秒でTimeoutするcontextを生成する
	// cancelを実行することでTimeout前にキャンセルを実行することができる
	//
	// また後述するようにGo1.7ではnet/httpパッケージでcontext
	// を扱えるようになる．例えば*http.Requestからそのリクエストの
	// contextを取得できる．
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	key = "foo"
	c := SetToken(ctx, key, "www")
	fmt.Printf("token is %v \n", c.Value(key))

	errCh := make(chan error, 1)
	go func() {
		errCh <- request(c)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			log.Println("failed:", err)
			return
		}
	}

	log.Println("success")
}

func request(ctx context.Context) error {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", "http://localhost:8081/ping", nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	// 新たにgoroutineを生成して実際のリクエストを行う
	// 結果はerror channelに投げる
	errCh := make(chan error, 1)
	go func() {
		_, err := client.Do(req)
		errCh <- err
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}

	// Timeoutが発生する，もしくはCancelが実行されると
	// Channelが返る
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-errCh
		return ctx.Err()
	}

	return nil
}

// SetToken - setter - set token to value.
func SetToken(ctx context.Context, key TokenKey, val string) context.Context {
	c := context.WithValue(ctx, key, val)
	fmt.Printf("token is %v \n", c.Value(key))
	return c
}

// GetToken - getter - get token to value.
func GetToken(ctx context.Context, key TokenKey) (string, error) {
	v := ctx.Value(key)

	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("token not found")
	}

	return token, nil
}
