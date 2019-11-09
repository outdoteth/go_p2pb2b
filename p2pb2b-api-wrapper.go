package go_p2pb2b

import (
	"context"
	"net/http"
)

//Client object for initial parameters
type Client struct {
	URL string
	API_Key string
	Ctx context.Context
}


//Initialiser function
func new_client(url, api_key string, ctx context.Context) *Client {
	if ctx == nil {
		ctx = context.Background()
	}

	return &Client {
		URL: url,
		API_Key: api_key,
		Ctx: ctx,
	}
}

func (clt *Client) API_request(method, endpoint, string) (*http.Response, err) {
	ctx := clt.Ctx
	if method == http.MethodGet {
		req, err := http.newRequest(method, clt.URL + endpoint, nil)
	}

	res, err := http.Get()
	if err != nil {
		return (res, err)
	}

	return (res, nil)
}

//this is where we need a custom response type so its easy for the end user
func (clt *Client) get_markets() (*http.Response, err) {
	endpoint := "/public/markets"
	res, err := clt.API_request(http.methodGet, endpoint)
	if err != nil {
		return (res, err)
	}
	return (res, nil)
}





