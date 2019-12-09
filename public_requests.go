package p2pb2b

import (
	"encoding/json"
	"fmMt"
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

type MarketsJsonRes struct {
	Success bool            `json: "success"`
	Message string          `json: "message"`
	Result  []marketsResult `json: "result"`
}

func (clt *Client) Markets() (*MarketsJsonRes, error) {
	endpoint := "api/v1/public/markets"
	res, err := clt.APIRequest(http.MethodGet, endpoint)

	if err != nil {
		return nil, err
	}

	var jsonRes MarketsJsonRes
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

type TickersJsonRes struct {
	Success     bool                     `json: "success"`
	Message     string                   `json: "message"`
	Result      map[string]tickersResult `json: "result"`
	CacheTime   float64                  `json: "cache_time"`
	CurrentTime float64                  `json: "current_time"`
}

func (clt *Client) Tickers() (*TickersJsonRes, error) {
	endpoint := "api/v1/public/tickers"
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes TickersJsonRes
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

type TickerJsonRes struct {
	Success     bool         `json: "success"`
	Message     string       `json: "message"`
	Result      tickerResult `json: "result"`
	CacheTime   float64      `json: "cache_time"`
	CurrentTime float64      `json: "current_time"`
}

type TickerParams struct {
	Symbol string
}

func (clt *Client) Ticker(opts TickerParams) (*TickerJsonRes, error) {
	endpoint := "api/v1/public/ticker?market=" + opts.Symbol
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes TickerJsonRes
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

type OrderBookJsonRes struct {
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

func (clt *Client) OrderBook(opts OrderBookParams) (*OrderBookJsonRes, error) {
	endpoint := "api/v1/public/book?market=" + opts.Market + "&side=" + opts.Side + "&offset=" + opts.Offset + "&limit=" + opts.Limit
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes OrderBookJsonRes
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

type HistoryJsonRes struct {
	Success     bool            `json: "success"`
	Message     string          `json: "message"`
	Result      []historyResult `json: "result"`
	CacheTime   float64         `json: "cache_time"`
	CurrentTime float64         `json: "current_time"`
}

type HistoryParams struct {
	Market  string
	LastId string
	Limit   string
}

func (clt *Client) History(opts HistoryParams) (*HistoryJsonRes, error) {
	endpoint := "api/v1/public/history?market=" + opts.Market + "&lastId=" + opts.LastId + "&limit=" + opts.Limit
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes HistoryJsonRes
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}

type depthResult struct {
	Asks [][2]string `json:"asks"`
	Bids [][2]string `json:"bids"`
}

type DepthJsonRes struct {
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

func (clt *Client) Depth(opts Depth_params) (*DepthJsonRes, error) {
	endpoint := "api/v1/public/depth/result?market=" + opts.Market + "&limit=" + opts.Limit
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes DepthJsonRes
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}

type productsResult struct {
	Id          string `json:"id"`
	FromSymbol string `json:"from_symbol"`
	ToSymbol   string `json:"to_symbol"`
}

type ProductsJsonRes struct {
	Success     bool             `json:"success"`
	Message     string           `json:"message"`
	Result      []productsResult `json:"result"`
	CacheTime   float64          `json:"cache_time"`
	CurrentTime float64          `json:"current_time"`
}

func (clt *Client) Products() (*ProductsJsonRes, error) {
	endpoint := "api/v1/public/products"
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes ProductsJsonRes
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil

}

type SymbolsJsonRes struct {
	Success     bool     `json:"success"`
	Message     string   `json:"message"`
	Result      []string `json:"result"`
	CacheTime   float64  `json:"cache_time"`
	CurrentTime float64  `json:"current_time"`
}

func (clt *Client) Symbols() (*SymbolsJsonRes, error) {
	endpoint := "api/v1/public/symbols"
	res, err := clt.APIRequest(http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes SymbolsJsonRes
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}
