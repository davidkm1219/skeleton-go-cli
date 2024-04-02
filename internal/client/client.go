// Package client provides the client for making HTTP requests.
package client

import (
	"context"
	"fmt"
	"net/http"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is a wrapper around the http client.
type Client struct {
	httpClient httpClient
}

// NewClient creates a new Client.
func NewClient(httpClient httpClient) *Client {
	return &Client{httpClient: httpClient}
}

// Get performs a GET request.
func (c *Client) Get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}

	return resp, nil
}
