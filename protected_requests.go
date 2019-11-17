package p2pb2b

import (
	"encoding/json"
	"net/http"
	"time"
)

type currencyResult struct {
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
}

type CurrencyJsonRes struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Result  currencyResult `json:"result"`
}

type CurrencyJsonBody struct {
	RequestUrl string `json:"request"`
	Nonce      int64  `json:"nonce"`
	Currency   string `json:"currency"`
}

type CurrencyBalanceParams struct {
	Currency string /// e.g "ETH"
}

func (clt *Client) CurrencyBalance(opts CurrencyBalanceParams) (*CurrencyJsonRes, error) {
	endpoint := "/api/v1/account/balance"
	postBody := CurrencyJsonBody{
		RequestUrl: endpoint,
		Nonce:      time.Now().Unix(),
		Currency:   opts.Currency,
	}

	res, err := clt.AuthAPIRequest(postBody, http.MethodPost, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes CurrencyJsonRes
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}
