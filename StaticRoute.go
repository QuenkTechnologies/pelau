package pelau

import (
	"net/http"
)

//StaticRoute provides static routing support.
type staticRoute struct {
	path   string
	action Callback
}

//Query looks up whether the current http path matches the static route..
func (s *staticRoute) Query(route string, w http.ResponseWriter, r *http.Request) bool {

	if s.path == route {

		s.action(DefaultRequest(r, make([]string, 0)), DefaultResponse(w))

		return true

	}

	return false

}

//Static Route constructor.
func Static(path string, callback Callback) Route {

	return &staticRoute{path, callback}

}
