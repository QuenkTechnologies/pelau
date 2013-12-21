package pelau

import (
	"net/http"
)

//Route is an inteface used to setup routing.
type Route interface {
	Query(string, http.ResponseWriter, *http.Request) bool
}
