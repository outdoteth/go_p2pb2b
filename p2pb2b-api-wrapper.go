package go_p2pb2b

//This needs some info and is the initialiser object- Perhaps API keys etc.
type p2pb2b_obj struct {}

//This needs to return a json object
func (e *p2pb2b_obj) get_price() float64 {
	return 5.0
}
