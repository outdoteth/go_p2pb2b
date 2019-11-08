package go_p2pb2b

import (
	"context"
)

//Client object for initial parameters
type Client struct {
	URL string
	API_Key string
	Ctx context.Context
}

//Initialiser function
func NewClient(url, api_key string, ctx context.Context) *Client {
	if ctx == nil {
		ctx = context.Background()
	}

	return &Client {
		URL: url,
		API_Key: api_key,
		Ctx: ctx,
	}
}

func (clt *Client) get_price() float64 {
	return 5.0
}
