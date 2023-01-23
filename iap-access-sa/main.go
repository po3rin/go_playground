package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/api/idtoken"
)

// makeIAPRequest makes a request to an application protected by Identity-Aware
// Proxy with the given audience.
func makeIAPRequest(w io.Writer, request *http.Request, audience string) error {
	// request, err := http.NewRequest("GET", "http://example.com", nil)
	// audience := "IAP_CLIENT_ID.apps.googleusercontent.com"
	ctx := context.Background()

	// client is a http.Client that automatically adds an "Authorization" header
	// to any requests made.
	client, err := idtoken.NewClient(ctx, audience)
	if err != nil {
		return fmt.Errorf("idtoken.NewClient: %v", err)
	}

	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("client.Do: %v", err)
	}
	defer response.Body.Close()
	if _, err := io.Copy(w, response.Body); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	return nil
}
