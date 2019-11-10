package go_p2pb2b

import (
	"context"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
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

func (clt *Client) API_request(method, endpoint string) ([]byte, error) {
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

type get_markets_result struct {
	Name string `json: "name"`
	Stock string `json: "stock"`
	Money string `json: "money"`
	MoneyPrec string `json: moneyPrec`
	StockPrec string `json: "stockPrec"`
	FeePrec string `json: "feePrec"`
	MinAmount string `json: "minAmount"`
}

type get_markets_json struct {
	Success bool `json: "success"`
	Message string `json: "message"`
	Result []get_markets_result `json: "result"`
}

func (clt *Client) get_markets() (*get_markets_json, error) {
	endpoint := "/public/markets"
	res, err := clt.API_request(http.MethodGet, endpoint)

	if err != nil {
		return nil, err
	}

	var json_res get_markets_json
	if err := json.Unmarshal(res, &json_res); err != nil {
		return nil, err
	}
	return &json_res, nil
}





