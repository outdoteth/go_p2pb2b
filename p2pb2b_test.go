package p2pb2b

import (
	"context"
	"testing"
)

/*
func Test_get_markets_get_request(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io/api/v1", "", ctx)
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
	client := NewClient("https://api.p2pb2b.io/api/v1", "", ctx)
	res, err := client.get_tickers()
	if err != nil {
		t.Errorf("market get request failed, expected %v, got %v", nil, err)
		return
	}

	t.Logf("market get request success, expected %+v\n, got %+v\n", res, res)
}*/

/*func Test_get_ticker_get_request(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io/api/v1", "", ctx)
	res, err := client.Get_ticker(Get_ticker_params{Symbol: "ETH_BTC"})
	if err != nil {
		t.Errorf("get_ticker get request failed, expected %v, got %v", nil, err)
		return
	}

	t.Logf("get_ticker get request success, expected %+v\n, got %+v\n", res, res)

}*/

/*func Test_order_book_get_request(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io/api/v1", "", ctx)
	res, err := client.OrderBook(OrderBook_params{Market: "ETH_BTC", Side: "buy", Offset: "0", Limit: "100"})
	if err != nil {
		t.Errorf("OrderBook get request failed, expected %v, got %v\n", nil, err)
	}

	t.Logf("OrderBook get request success, expected %+v\n, got %v+\n", res, res)
}*/

/*func Test_history_get_request(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io/api/v1", "", ctx)
	res, err := client.History(History_params{Market: "ETH_BTC", Last_id: "1",  Limit: "100"})
	if err != nil {
		t.Errorf("History get request failed, expected %v, got %v\n", nil, err)
	}

	if !res.Success {
		t.Errorf("History get request failed, expected res.Success to be true,instead got %+v\n", res)
	}

	if len(res.Result) == 0 {
		t.Errorf("History get request failed, expected res.Result length to be greater than 1, instead got %+v\n", res)
		return
	}

	if i := res.Result[0]; i.Id == 0 || i.Type == "" || i.Time == 0.0 || i.Amount == "" || i.Price == "" {
		t.Errorf("History get request failed, expected res.Result[0] to have values in all fields, instead got %+v\n", res)
		return
	}

	t.Logf("History get request success")
}*/

/*func TestDepth(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io/api/v1", "", ctx)
	res, err := client.Depth(Depth_params{Market: "ETH_BTC", Limit: "100"})
	if err != nil {
		t.Errorf("OrderBook get request failed, expected %v, got %v\n", nil, err)
	}

	if !res.Success {
		t.Errorf("Depth() get request failed, expected res.Success to be true,instead got %+v\n", res)
		return
	}

	if len(res.Result.Asks) < 1 || len(res.Result.Bids) < 1 {
		t.Errorf("Depth() get request failed, expected to get some Asks or Bids, instead got %+v\n", res)
		return
	}

	if len(res.Result.Asks[0]) != 2 || len(res.Result.Bids[0]) != 2 {
		t.Errorf("Depth() get request failed. res.Result.Bids or res.Result.Asks has malformed data, expected [][2]string, instead got %v+\n", res)
	}
	t.Logf("OrderBook get request success, expected %+v\n, got %v+\n", res, res)
}*/

/*func TestProducts(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io/api/v1", "", ctx)
	res, err := client.Products()
	if err != nil {
		t.Errorf("Products() get request failed, expected Error=%v, instead got Error=%v\n", nil, err)
		return
	}

	if !res.Success {
		t.Errorf("Products() get request failed, expected res.Success to be true, instead got %+v\n", res)
		return
	}

	if len(res.Result) < 1 {
		t.Errorf("Products() get request failed, expected res.Result length to have data, instead got %+v\n", res)
		return
	}

	t.Logf("Products() get request success, expected %+v\n, got %v+\n", res, res)
}*/

func TestSymbols(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io/api/v1", "", ctx)
	res, err := client.Symbols()
	if err != nil {
		t.Errorf("Symbols() get request failed, expected Error=%v, instead got Error=%v\n", nil, err)
		return
	}

	if !res.Success {
		t.Errorf("Symbols() get request failed, expected res.Success to be true, instead got %+v\n", res)
		return
	}

	if len(res.Result) < 1 {
		t.Errorf("Symbols() get request failed, expected res.Result length to have data, instead got %+v\n", res)
		return
	}

	t.Logf("Symbol() get request success, expected %+v\n, got %v+\n", res, res)
}
