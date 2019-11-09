package go_p2pb2b

import (
	"context"
	"net/http"
	"io/ioutil"
)

//Client object for initial parameters
type Client struct {
	URL string
	API_Key string
	Ctx context.Context
}


//Initialiser function
func New_Client(url, api_key string, ctx context.Context) *Client {
	if ctx == nil {
		ctx = context.Background()
	}

	return &Client {
		URL: url,
		API_Key: api_key,
		Ctx: ctx,
	}
}

func (clt *Client) API_request(method, endpoint string) (*http.Response, error) {
	ctx := clt.Ctx
	var req *http.Request
	var err error
	if method == http.MethodGet {
		req, err = http.NewRequest(method, clt.URL + endpoint, nil)
	}

	if method == http.MethodPost {
		//run through all of the signing procedures
	}

	req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return res, err
	}

	return res, nil
}



//this is where we need a custom response type so its easy for the end user
func (clt *Client) get_markets() (string, error) {
	endpoint := "/public/markets"
	res, err := clt.API_request(http.MethodGet, endpoint)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}





