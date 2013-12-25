package pelau

import ()

//Route is an inteface used to setup routing.
type Route interface {
	Query(string, Request, Response) bool
}
