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

type Get_markets_json struct {
	Success bool `json: "success"`
	Message string `json: "message"`
	Result []get_markets_result `json: "result"`
}

func (clt *Client) Get_markets() (*Get_markets_json, error) {
	endpoint := "/public/markets"
	res, err := clt.API_request(http.MethodGet, endpoint)

	if err != nil {
		return nil, err
	}

	var json_res Get_markets_json
	if err := json.Unmarshal(res, &json_res); err != nil {
		return nil, err
	}
	return &json_res, nil
}

type get_tickers_result struct {
	At int `json: "at"`
	Ticker struct {
		Bid string `json: "bid"`
		Ask string `json: "ask"`
		Low  string `json: "low"`
		High string `json: "high"`
		Last string `json: "last"`
		Vol string `json: "vol"`
		Change string `json: "change"`
	} `json: "ticker"`
}

type Get_tickers_json struct {
	Success bool `json: "success"`
	Message string `json: "message"`
	Result map[string]get_tickers_result `json: "result"`
	Cache_time float64 `json: "cache_time"`
	Current_time float64 `json: "current_time"`
}

func (clt *Client) Get_tickers() (*Get_tickers_json, error) {
	endpoint := "/public/tickers"
	res, err := clt.API_request(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(res))
	var json_res Get_tickers_json
	if err := json.Unmarshal(res, &json_res); err != nil {
		return nil, err
	}

	return &json_res, nil
}

type get_ticker_result struct {
	Bid string `json: "bid"`
	Ask string `json: "ask"`
	Open string `json: "open"`
	High string `json: "high"`
	Low  string `json: "low"`
	Last string `json: "last"`
	Volume string `json: "volume"`
	Deal string `json: "deal"`
	Change string `json: "change"`
}

type Get_ticker_json struct {
	Success bool `json: "success"`
	Message string `json: "message"`
	Result get_ticker_result `json: "result"`
	Cache_time float64 `json: "cache_time"`
	Current_time float64 `json: "current_time"`
}
type Get_ticker_params struct {
	Symbol string
}

func (clt *Client) Get_ticker(opts Get_ticker_params) (*Get_ticker_json, error) {
	endpoint := "/public/ticker?market=" + opts.Symbol
	res, err := clt.API_request(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}
	fmt.Println("\n" + string(res))
	var json_res Get_ticker_json
	if err := json.Unmarshal(res, &json_res); err != nil {
		return nil, err
	}

	return &json_res, nil
}
