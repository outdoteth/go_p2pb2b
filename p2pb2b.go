package p2pb2b

import (
	"context"
)

//Client object for initial parameters
type Client struct {
	URL    string
	APIKey string
	Ctx    context.Context
}

//Initialiser function
func NewClient(url, apiKey string, ctx context.Context) *Client {
	if ctx == nil {
		ctx = context.Background()
	}

	return &Client{
		URL:    url,
		APIKey: apiKey,
		Ctx:    ctx,
	}
}
