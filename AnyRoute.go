package pelau

import (
	"net/http"
)

type anyRoute struct {
	action Callback
}

//Query implements from the Route interface
func (r anyRoute) Query(path string, res http.ResponseWriter, req *http.Request) bool {

	r.action(DefaultRequest(req, make([]string, 0)), DefaultResponse(res))

	return true

}

//Any creates a new AnyRoute
func Any(cb Callback) Route {

	return &anyRoute{cb}

}
