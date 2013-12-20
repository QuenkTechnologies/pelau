package pelau

//Route is an inteface used to setup routing.
type Route interface {
	Query(string) bool
	Execute(*Request, *Response)
}
