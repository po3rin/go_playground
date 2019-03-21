package openapi_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/po3rin/gooatest"
	openapi "github.com/po3rin/userapi/go"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	// テストケースを定義します。今回は1パターン。
	tests := []struct {
		name       string
		input      string
		want       string
		wantStatus int
	}{
		{
			name:       "get user",
			want:       `{"users":[{"id":1,"name":"po3rin"},{"id":2,"name":"po4rin"}]}`,
			wantStatus: 200,
		},
	}

	// テストケースの数だけforで回します。
	for _, tt := range tests {
		httpReq, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Fatalf("test %v unexpected error : %v", tt.name, err)
		}
		httpReq.Header.Set("Content-Type", "application/json; charset=utf-8")

		// httptestパッケージを使ってHandlerからレスポンスを所得します。
		r := openapi.NewRouter()
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httpReq)

		// assertion
		assert.Equal(t, tt.wantStatus, w.Code)
		assert.Equal(t, tt.want, w.Body.String())

		// paramsのにvalidateに必要な設定を渡す
		p := gooatest.Params{
			HTTPReq:    httpReq,
			BaseURL:    "http://localhost:8080",
			SchemaPath: "./../../schema/schema.yml",
			Context:    context.Background(),
			HTTPRes:    w.Result(),
		}

		// paramsを使ったvalidator初期化
		v, err := gooatest.NewValidator(p)
		if err != nil {
			t.Fatalf("unexpected error : %v", err)
		}

		// validate実行
		err = v.ValidateRequest()
		if err != nil {
			t.Fatalf("unexpected error : %v", err)
		}
		err = v.ValidateResponse()
		if err != nil {
			t.Fatalf("unexpected error : %v", err)
		}
	}
}
