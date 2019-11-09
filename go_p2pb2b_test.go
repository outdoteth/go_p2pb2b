package go_p2pb2b

import(
	"testing"
	"context"
	"fmt"
)

func Test_market_request(t *testing.T) {
	ctx := context.Background()
	client := New_Client("https://api.p2pb2b.io/api/v1", "", ctx)
	res, err := client.get_markets()
	fmt.Print("hi")
	if err != nil {
		t.Errorf("market get request failed, expected %v, got %v", nil, err)
	}

	t.Logf("market get request success, expected %v, got %v", res, res)
}
