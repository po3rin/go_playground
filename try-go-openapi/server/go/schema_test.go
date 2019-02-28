package openapi_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/po3rin/go-playground/try-go-openapi/server/go"
)

func NewRouterFromYAML(path string) (*openapi3filter.Router, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromYAMLData(data)
	if err != nil {
		return nil, err
	}
	router := openapi3filter.NewRouter()
	router.AddSwagger(swagger)
	if err != nil {
		return nil, err
	}
	return router, nil
}

func TestCreatePost(t *testing.T) {
	tests := []struct {
		input openapi.Post
	}{
		{
			input: openapi.Post{
				Post: &openapi.PostProperties{
					Title:   "hello",
					Content: "world",
				},
			},
		},
		{
			input: openapi.Post{
				Post: &openapi.PostProperties{
					Title:   "",
					Content: "",
				},
			},
		},
	}

	var (
		respStatus      = 200
		respContentType = "application/json"
		respBody        = bytes.NewBufferString(`{}`)
	)

	router, err := NewRouterFromYAML("./../../schema.yml")
	if err != nil {
		t.Fatalf("unexpected error : %v", err)
	}

	ctx := context.Background()
	for _, tt := range tests {
		d, err := json.Marshal(tt.input)
		if err != nil {
			t.Fatalf("unexpected error : %v", err)
		}
		reqBody := bytes.NewBuffer(d)

		httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/posts", reqBody)
		if err != nil {
			t.Fatalf("unexpected error : %v", err)
		}
		httpReq.Header.Set("Content-Type", "application/json; charset=utf-8")
		route, _, err := router.FindRoute(httpReq.Method, httpReq.URL)
		if err != nil {
			t.Fatalf("unexpected error : %v", err)
		}
		// Validate request
		requestValidationInput := &openapi3filter.RequestValidationInput{
			Request: httpReq,
			Route:   route,
		}
		if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
			t.Fatalf("unexpected error : %v", err)
		}
		responseValidationInput := &openapi3filter.ResponseValidationInput{
			RequestValidationInput: requestValidationInput,
			Status:                 respStatus,
			Header: http.Header{
				"Content-Type": []string{
					respContentType,
				},
			},
		}
		if respBody != nil {
			data, err := json.Marshal(respBody)
			if err != nil {
				t.Fatalf("unexpected error : %v", err)
			}
			t.Log(string(data))
			responseValidationInput.SetBodyBytes(data)
		}
		if err := openapi3filter.ValidateResponse(ctx, responseValidationInput); err != nil {
			t.Fatalf("unexpected error : %v", err)
		}
	}
}
