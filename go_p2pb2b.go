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

type markets_result struct {
	Name string `json: "name"`
	Stock string `json: "stock"`
	Money string `json: "money"`
	MoneyPrec string `json: moneyPrec`
	StockPrec string `json: "stockPrec"`
	FeePrec string `json: "feePrec"`
	MinAmount string `json: "minAmount"`
}

type Markets_json struct {
	Success bool `json: "success"`
	Message string `json: "message"`
	Result []markets_result `json: "result"`
}

func (clt *Client) Markets() (*Markets_json, error) {
	endpoint := "/public/markets"
	res, err := clt.API_request(http.MethodGet, endpoint)

	if err != nil {
		return nil, err
	}

	var json_res Markets_json
	if err := json.Unmarshal(res, &json_res); err != nil {
		return nil, err
	}
	return &json_res, nil
}

type tickers_result struct {
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

type Tickers_json struct {
	Success bool `json: "success"`
	Message string `json: "message"`
	Result map[string]tickers_result `json: "result"`
	Cache_time float64 `json: "cache_time"`
	Current_time float64 `json: "current_time"`
}

func (clt *Client) Tickers() (*Tickers_json, error) {
	endpoint := "/public/tickers"
	res, err := clt.API_request(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(res))
	var json_res Tickers_json
	if err := json.Unmarshal(res, &json_res); err != nil {
		return nil, err
	}

	return &json_res, nil
}

type ticker_result struct {
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

type Ticker_json struct {
	Success bool `json: "success"`
	Message string `json: "message"`
	Result ticker_result `json: "result"`
	Cache_time float64 `json: "cache_time"`
	Current_time float64 `json: "current_time"`
}

type Ticker_params struct {
	Symbol string
}

func (clt *Client) Ticker(opts Ticker_params) (*Ticker_json, error) {
	endpoint := "/public/ticker?market=" + opts.Symbol
	res, err := clt.API_request(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}
	fmt.Println("\n" + string(res))
	var json_res Ticker_json
	if err := json.Unmarshal(res, &json_res); err != nil {
		return nil, err
	}

	return &json_res, nil
}


type order_book_result struct {
	Offset int64 `json: "offset"`
	Limit int64 `json: "limit"`
	Total int64 `json: "total"`
	Orders []struct {
		Id int64 `json: "id"`
		Left string `json: "left"`
		Market string `json: "market"`
		Amount string `json: "amount"`
		Type string `json: "type"`
		Price string `json: "price"`
		Timestamp float64 `json: "timestamp"`
		Side string `json: "side"`
		DealFee string `json: "dealFee"`
		TakerFee string `json: "takerFee"`
		MakerFee string `json: "makerFee"`
		DealStock string `json: "dealStock"`
		DealMoney string `json: "dealMoney"`
	} `json: "orders"`
}

type Order_book_json struct {
	Success bool `json: "success"`
	Message string `json: "message"`
	Result order_book_result `json: "result"`
	Cache_time float64 `json: "cache_time"`
	Current_time float64 `json: "current_time"`
}

type Order_book_params struct {
	Market string
	Side string
	Offset string
	Limit string
}

func (clt *Client) Order_book(opts Order_book_params) (*Order_book_json, error) {
	endpoint := "/public/book?market=" + opts.Market + "&side=" + opts.Side + "&offset=" + opts.Offset + "&limit=" + opts.Limit
	res, err := clt.API_request(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	fmt.Println("\n" + string(res))
	var json_res Order_book_json
	if err := json.Unmarshal(res, &json_res); err != nil {
		return nil, err
	}

	return &json_res, nil
}

type history_result struct {
	Id int64 `json: "id"`
	Type string `json: "type"`
	Time float64 `json: "time"`
	Amount string `json: "amount"`
	Price string `json: "price"`
}

type History_json struct {
	Success bool `json: "success"`
	Message string `json: "message"`
	Result []history_result `json: "result"`
	Cache_time float64 `json: "cache_time"`
	Current_time float64 `json: "current_time"`
}

type History_params struct {
	Market string
	Last_id string
	Limit string
}

func (clt *Client) History(opts History_params) (*History_json, error) {
	endpoint := "/public/history?market=" + opts.Market + "&lastId=" + opts.Last_id + "&limit=" + opts.Limit
	res, err := clt.API_request(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	fmt.Println("\n" + string(res))
	var json_res History_json
	if err := json.Unmarshal(res, &json_res); err != nil {
		return nil, err
	}

	return &json_res, nil
}


