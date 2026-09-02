package ntfy

import (
	"net/http"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func New(url string) *Client {
	return &Client{
		baseURL: url,
		client:  http.DefaultClient,
	}
}
