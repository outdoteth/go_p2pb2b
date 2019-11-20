package p2pb2b

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type marketsResult struct {
	Name      string `json: "name"`
	Stock     string `json: "stock"`
	Money     string `json: "money"`
	MoneyPrec string `json: moneyPrec`
	StockPrec string `json: "stockPrec"`
	FeePrec   string `json: "feePrec"`
	MinAmount string `json: "minAmount"`
}

type MarketsJson struct {
	Success bool            `json: "success"`
	Message string          `json: "message"`
	Result  []marketsResult `json: "result"`
}

func (clt *Client) Markets() (*MarketsJson, error) {
	endpoint := "/public/markets"
	res, err := clt.APIRequest(http.MethodGet, endpoint)

	if err != nil {
		return nil, err
	}

	var jsonRes MarketsJson
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}
	return &jsonRes, nil
}

type tickersResult struct {
	At     int `json: "at"`
	Ticker struct {
		Bid    string `json: "bid"`
		Ask    string `json: "ask"`
		Low    string `json: "low"`
		High   string `json: "high"`
		Last   string `json: "last"`
		Vol    string `json: "vol"`
		Change string `json: "change"`
	} `json: "ticker"`
}

type TickersJson struct {
	Success     bool                     `json: "success"`
	Message     string                   `json: "message"`
	Result      map[string]tickersResult `json: "result"`
	CacheTime   float64                  `json: "cache_time"`
	CurrentTime float64                  `json: "current_time"`
}

func (clt *Client) Tickers() (*TickersJson, error) {
	endpoint := "/public/tickers"
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(res))
	var jsonRes TickersJson
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}

type tickerResult struct {
	Bid    string `json: "bid"`
	Ask    string `json: "ask"`
	Open   string `json: "open"`
	High   string `json: "high"`
	Low    string `json: "low"`
	Last   string `json: "last"`
	Volume string `json: "volume"`
	Deal   string `json: "deal"`
	Change string `json: "change"`
}

type TickerJson struct {
	Success     bool         `json: "success"`
	Message     string       `json: "message"`
	Result      tickerResult `json: "result"`
	CacheTime   float64      `json: "cache_time"`
	CurrentTime float64      `json: "current_time"`
}

type TickerParams struct {
	Symbol string
}

func (clt *Client) Ticker(opts TickerParams) (*TickerJson, error) {
	endpoint := "/public/ticker?market=" + opts.Symbol
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}
	fmt.Println("\n" + string(res))
	var jsonRes TickerJson
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}

type orderBookResult struct {
	Offset uint64 `json: "offset"`
	Limit  uint64 `json: "limit"`
	Total  uint64 `json: "total"`
	Orders []struct {
		Id        uint64  `json: "id"`
		Left      string  `json: "left"`
		Market    string  `json: "market"`
		Amount    string  `json: "amount"`
		Type      string  `json: "type"`
		Price     string  `json: "price"`
		Timestamp float64 `json: "timestamp"`
		Side      string  `json: "side"`
		DealFee   string  `json: "dealFee"`
		TakerFee  string  `json: "takerFee"`
		MakerFee  string  `json: "makerFee"`
		DealStock string  `json: "dealStock"`
		DealMoney string  `json: "dealMoney"`
	} `json: "orders"`
}

type OrderBookJson struct {
	Success     bool            `json: "success"`
	Message     string          `json: "message"`
	Result      orderBookResult `json: "result"`
	CacheTime   float64         `json: "cache_time"`
	CurrentTime float64         `json: "current_time"`
}

type OrderBookParams struct {
	Market string
	Side   string
	Offset string
	Limit  string
}

func (clt *Client) OrderBook(opts OrderBookParams) (*OrderBookJson, error) {
	endpoint := "/public/book?market=" + opts.Market + "&side=" + opts.Side + "&offset=" + opts.Offset + "&limit=" + opts.Limit
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	fmt.Println("\n" + string(res))
	var jsonRes OrderBookJson
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}

type historyResult struct {
	Id     uint64  `json: "id"`
	Type   string  `json: "type"`
	Time   float64 `json: "time"`
	Amount string  `json: "amount"`
	Price  string  `json: "price"`
}

type HistoryJson struct {
	Success     bool            `json: "success"`
	Message     string          `json: "message"`
	Result      []historyResult `json: "result"`
	CacheTime   float64         `json: "cache_time"`
	CurrentTime float64         `json: "current_time"`
}

type HistoryParams struct {
	Market  string
	Last_id string
	Limit   string
}

func (clt *Client) History(opts HistoryParams) (*HistoryJson, error) {
	endpoint := "/public/history?market=" + opts.Market + "&lastId=" + opts.Last_id + "&limit=" + opts.Limit
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	fmt.Println("\n" + string(res))
	var jsonRes HistoryJson
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}

type depthResult struct {
	Asks [][2]string `json:"asks"`
	Bids [][2]string `json:"bids"`
}

type DepthJson struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	Result      depthResult `json:"result"`
	CacheTime   float64     `json:"cache_time"`
	CurrentTime float64     `json:"current_time"`
}

type Depth_params struct {
	Market string
	Limit  string
}

func (clt *Client) Depth(opts Depth_params) (*DepthJson, error) {
	endpoint := "/public/depth/result?market=" + opts.Market + "&limit=" + opts.Limit
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	fmt.Println("\n" + string(res))
	var jsonRes DepthJson
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}

type productsResult struct {
	Id          string `json:"id"`
	From_symbol string `json:"from_symbol"`
	To_symbol   string `json:"to_symbol"`
}

type ProductsJson struct {
	Success     bool             `json:"success"`
	Message     string           `json:"message"`
	Result      []productsResult `json:"result"`
	CacheTime   float64          `json:"cache_time"`
	CurrentTime float64          `json:"current_time"`
}

func (clt *Client) Products() (*ProductsJson, error) {
	endpoint := "/public/products"
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	fmt.Println("\n" + string(res))
	var jsonRes ProductsJson
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil

}

type SymbolsJson struct {
	Success     bool     `json:"success"`
	Message     string   `json:"message"`
	Result      []string `json:"result"`
	CacheTime   float64  `json:"cache_time"`
	CurrentTime float64  `json:"current_time"`
}

func (clt *Client) Symbols() (*SymbolsJson, error) {
	endpoint := "/public/symbols"
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	fmt.Println("\n" + string(res))
	var jsonRes SymbolsJson
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}
