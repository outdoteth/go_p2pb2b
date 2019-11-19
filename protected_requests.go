package p2pb2b

import (
	"encoding/json"
	"net/http"
	"time"
)

type currencyBalanceResult struct {
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
}

type CurrencyBalanceJsonRes struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Result  currencyBalanceResult `json:"result"`
}

type CurrencyBalanceJsonBody struct {
	RequestUrl string `json:"request"`
	Nonce      int64  `json:"nonce"`
	Currency   string `json:"currency"`
}

type CurrencyBalanceParams struct {
	Currency string /// e.g "ETH"
}

func (clt *Client) CurrencyBalance(opts CurrencyBalanceParams) (*CurrencyBalanceJsonRes, error) {
	endpoint := "/api/v1/account/balance"
	postBody := CurrencyBalanceJsonBody{
		RequestUrl: endpoint,
		Nonce:      time.Now().UnixNano(),
		Currency:   opts.Currency,
	}

	res, err := clt.AuthAPIRequest(postBody, http.MethodPost, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes CurrencyBalanceJsonRes
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}

type BalancesJsonBody struct {
	RequestUrl string `json:"request"`
	Nonce      int64  `json:"nonce"`
}

type balancesResult struct {
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
}

type BalancesJsonRes struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Result  map[string]balancesResult `json:"result"`
}

func (clt *Client) Balances() (*BalancesJsonRes, error) {
	endpoint := "/api/v1/account/balances"
	postBody := BalancesJsonBody{
		RequestUrl: endpoint,
		Nonce:      time.Now().UnixNano(),
	}

	res, err := clt.AuthAPIRequest(postBody, http.MethodPost, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes BalancesJsonRes
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}

type CreateOrderParams struct {
	Market string `json:"market"`
	Side   string `json:"side"`
	Amount string `json:"amount"`
	Price  string `json:"price"`
}

type CreateOrderJsonBody struct {
	RequestUrl string `json:"request"`
	Nonce      int64  `json:"nonce"`
	Market     string `json:"market"`
	Side       string `json:"side"`
	Amount     string `json:"amount"`
	Price      string `json:"price"`
}

type createOrderResult struct {
	OrderId   int64   `json:"orderId"`
	Market    string  `json:"market"`
	Price     string  `json:"price"`
	Side      string  `json:"side"`
	Timestamp float64 `json:"timestamp"`
	DealMoney string  `json:"dealMoney"`
	DealStock string  `json:"dealStock"`
	Amount    string  `json:"amount"`
	TakerFee  string  `json:"takerFee"`
	MakerFee  string  `json:"makerFee"`
	Left      string  `json:"left"`
	DealFee   string  `json:"dealFee"`
}

type CreateOrderJsonRes struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Result  createOrderResult `json:"result"`
}

func (clt *Client) CreateOrder(opts CreateOrderParams) (*CreateOrderJsonRes, error) {
	endpoint := "/api/v1/order/new"
	postBody := CreateOrderJsonBody{
		RequestUrl: endpoint,
		Nonce:      time.Now().UnixNano(),
		Market:     opts.Market,
		Side:       opts.Side,
		Amount:     opts.Amount,
		Price:      opts.Price,
	}

	res, err := clt.AuthAPIRequest(postBody, http.MethodPost, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes CreateOrderJsonRes
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}

type cancelOrderResult struct {
	OrderId   int64   `json:"orderId"`
	Market    string  `json:"market"`
	Price     string  `json:"price"`
	Side      string  `json:"side"`
	Type      string  `json:"type"`
	Timestamp float64 `json:"timestamp"`
	DealMoney string  `json:"dealMoney"`
	DealStock string  `json:"dealStock"`
	Amount    string  `json:"amount"`
	TakerFee  string  `json:"takerFee"`
	MakerFee  string  `json:"makerFee"`
	Left      string  `json:"left"`
	DealFee   string  `json:"dealFee"`
}

type CancelOrderJsonRes struct {
	Success bool
	Message string
	Result  cancelOrderResult
}

type CancelOrderParams struct {
	Market  string
	OrderId int64
}

type CancelOrderJsonBody struct {
	RequestUrl string `json:"request"`
	Nonce      int64  `json:"nonce"`
	Market     string `json:"market"`
	OrderId    int64  `json:"orderId"`
}

func (clt *Client) CancelOrder(opts CancelOrderParams) (*CancelOrderJsonRes, error) {
	endpoint := "/api/v1/order/cancel"
	postBody := CancelOrderJsonBody{
		RequestUrl: endpoint,
		Nonce:      time.Now().UnixNano(),
		Market:     opts.Market,
		OrderId:    opts.OrderId,
	}

	res, err := clt.AuthAPIRequest(postBody, http.MethodPost, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes CancelOrderJsonRes
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil

}
