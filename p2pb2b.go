package p2pb2b

import (
	"context"
)

//Client object for initial parameters
type Client struct {
	URL       string
	APIKey    string
	APISecret string
	Ctx       context.Context
}

//Initialiser function
func NewClient(url, apiKey, apiSecret string, ctx context.Context) *Client {
	if ctx == nil {
		ctx = context.Background()
	}

	return &Client{
		URL:       url,
		APIKey:    apiKey,
		APISecret: apiSecret,
		Ctx:       ctx,
	}
}
