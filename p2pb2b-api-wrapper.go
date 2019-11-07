package go_p2pb2b

//This needs some info and is the initialiser object- Perhaps API keys etc.
type P2PB2B struct {}

//This needs to return a json object
func (e *P2PB2B) get_price() float64 {
	return 5.0
}
