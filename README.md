# GoP2PB2B
A lightweight Go wrapper around the p2pb2b api endpoints.

  **To Install:** 

	go get -v https://github.com/Dylan-Phoon/go_p2pb2b

**Example:**
```Go
import (
			"context"
			"fmt"
			"github.com/Dylan-Phoon/GoP2PB2B/p2pb2b"
		)
		
// These can be blank if you are not making protected requests
apiSecret := ""
apiKey := ""

ctx := context.Background()
client := p2pb2b.NewClient("https://api.p2pb2b.io", apiSecret, apiKey, ctx)

var res *MarketsJsonRes, err = client.Markets()
if err != nil {
	fmt.Printf("Markets() error: %v", err)
}
fmt.Printf("Markets() response: %+v", res)
```
**Each method follows the standard convention;**

```Go
func SomeRequest(SomeRequestParams) (*SomeRequestJsonRes, err) 	
```

## Contents
**Public Requests**
[Markets](#markets)
[Ticker](#ticker)
[Tickers](#tickers)
[OrderBook](#orderbook)
[History](#history)
[Depth](#depth)
[Product](#product)
[Symbols](#symbols)


### Markets 
Get markets info
Usage;
```Go
var res *MarketsJsonRes = client.Markets()
```

Returns;
```Go
type MarketsJsonRes struct {
	Success bool
	Message string
	Result []marketsResult {
				Name string
				Stock string
				Money string
				MoneyPrec string
				StockPrec string
				FeePrec string
				MinAmount string
			}
}
```


### Ticker 
Get info for a specified ticker
Usage;
```Go
params := TickerParams { 
	Symbol: "ETHBTC"
}
var res *TickerJsonRes = client.Ticker(params)
```
Parameters;
```Go
type TickerParams struct {
	Symbol string
}
```
Returns;
```Go
type TickerJsonRes struct {
	Success     bool         
	Message     string       
	Result      tickerResult {
					At     int 
					Ticker struct {
						Bid    string 
						Ask    string 
						Low    string 
						High   string 
						Last   string 
						Vol    string 
						Change string 
					} 
				}
	CacheTime   float64      
	CurrentTime float64      
}
```

### Tickers 
Get info for all tickers

Usage;
```Go
var res *TickersJsonRes = client.Tickers()
```
Returns;
```Go
type TickersJsonRes struct {
	Success     bool                     
	Message     string                   
	Result      map[string]tickersResult {
					At     int
					Ticker struct {
						Bid    string
						Ask    string
						Low    string
						High   string
						Last   string
						Vol    string
						Change string
					}
				}
	CacheTime   float64                  
	CurrentTime float64                  
}
```

### OrderBook 
Get info for an order book on a given ticker

Usage;
```Go
params := OrderBookParams {
	Market: "ETHBTC",
	Side: "buy",
	Offset: "0",
	Limit: "100"
}
var res *OrderBookJsonRes = client.OrderBook(params)
```
Parameters;
```Go
type OrderBookParams struct {
	Market string
	Side   string
	Offset string
	Limit  string
}
```
Returns;
```Go
type OrderBookJsonRes struct {
	Success     bool           
	Message     string         
	Result      orderBookResult {
					Offset uint64
					Limit  uint64
					Total  uint64
					Orders []struct {
						Id        uint64 
						Left      string 
						Market    string 
						Amount    string 
						Type      string 
						Price     string 
						Timestamp float64
						Side      string 
						DealFee   string 
						TakerFee  string 
						MakerFee  string 
						DealStock string 
						DealMoney string 
					}
				}
	CacheTime   float64        
	CurrentTime float64        
}
```
### History 
Get the last n trades up until a given trade id for a given ticker

Usage;
```Go
params := HistoryParams {
	Market: "ETHBTC",
	LastId: "1",
	Limit: "100"
}
var res *HistoryJsonRes = client.History(params)
```
Parameters;
```Go
type HistoryParams struct {
	Market string
	LastId   string
	Limit  string
}
```
Returns;
```Go
type HistoryJsonRes struct {
	Success     bool           
	Message     string         
	Result      []historyResult {
					Id     uint64 
					Type   string 
					Time   float64
					Amount string 
					Price  string 
				}
	CacheTime   float64        
	CurrentTime float64             
}
```

### Depth 
Get the depth of an order book for a given ticker

Usage;
```Go
params := DepthParams {
	Market "ETHBTC"
	Limit  "100"
}
var res *DepthJsonRes = client.Depth(params)
```
Parameters;
```Go
type DepthParams struct {
	Market string
	Limit  string
}
```
Returns;
```Go
type DepthJsonRes struct {
	Success     bool       
	Message     string     
	Result      depthResult {
					Asks [][2]string
					Bids [][2]string
				}
	CacheTime   float64    
	CurrentTime float64               
}
```

### Products 
Get all of the products

Usage;
```Go
var res *ProductsJsonRes = client.Products()
```
Returns;
```Go
type ProductsJsonRes struct {
	Success     bool             
	Message     string           
	Result      []productsResult {
					Id          string 
					FromSymbol 	string 
					ToSymbol   	string 
				}
	CacheTime   float64          
	CurrentTime float64                     
}
```

### Symbols 
Get all of the symbols

Usage;
```Go
var res *SymbolJsonRes = client.Symbol()
```
Returns;
```Go
type SymbolJsonRes struct {
	Success     bool    
	Message     string  
	Result      []string
	CacheTime   float64 
	CurrentTime float64                  
}
```