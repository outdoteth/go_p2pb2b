package go_p2pb2b

import(
	"testing"
	"context"
)

/*
func Test_get_markets_get_request(t *testing.T) {
	ctx := context.Background()
	client := New_Client("https://api.p2pb2b.io/api/v1", "", ctx)
	res, err := client.get_markets()
	if err != nil {
		t.Errorf("market get request failed, expected %v, got %v", nil, err)
		return
	}

	t.Logf("market get request success, expected %+v\n, got %+v\n", res.Success, res.Result[0])
}*/
/*
func Test_get_tickers_get_request(t *testing.T) {
	ctx := context.Background()
	client := New_Client("https://api.p2pb2b.io/api/v1", "", ctx)
	res, err := client.get_tickers()
	if err != nil {
		t.Errorf("market get request failed, expected %v, got %v", nil, err)
		return
	}

	t.Logf("market get request success, expected %+v\n, got %+v\n", res, res)
}*/

/*func Test_get_ticker_get_request(t *testing.T) {
	ctx := context.Background()
	client := New_Client("https://api.p2pb2b.io/api/v1", "", ctx)
	res, err := client.Get_ticker(Get_ticker_params{Symbol: "ETH_BTC"})
	if err != nil {
		t.Errorf("get_ticker get request failed, expected %v, got %v", nil, err)
		return
	}

	t.Logf("get_ticker get request success, expected %+v\n, got %+v\n", res, res)

}*/

func Test_order_book_get_request(t *testing.T) {
	ctx := context.Background()
	client := New_Client("https://api.p2pb2b.io/api/v1", "", ctx)
	res, err := client.Order_book(Order_book_params{Market: "ETH_BTC", Side: "buy", Offset: "0", Limit: "100"})
	if err != nil {
		t.Errorf("Order_book get request failed, expected %v, got %v\n", nil, err)
	}

	t.Logf("Order_book get request success, expected %+v\n, got %v+\n", res, res)
}



