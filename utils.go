package p2pb2b

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (clt *Client) APIRequest(method, endpoint string) ([]byte, error) {
	ctx := clt.Ctx
	var req *http.Request
	var err error
	if method == http.MethodGet {
		req, err = http.NewRequest(method, clt.URL+endpoint, nil)
	}

	if method == http.MethodPost {
		//run through all of the signing procedures
	}

	req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Print(string(body))
	return body, nil
}
