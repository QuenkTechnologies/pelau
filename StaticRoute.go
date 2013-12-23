package pelau

import ()

//StaticRoute provides static routing support.
type staticRoute struct {
	path   string
	action Callback
}

//Query looks up whether the current http path matches the static route..
func (s *staticRoute) Query(route string, req Request, res Response) bool {

	if s.path == route {

		req.SetParams(make([]string, 0))
		s.action(req, res)
		return true

	}

	return false

}

//Static Route constructor.
func Static(path string, callback Callback) Route {

	return &staticRoute{path, callback}

}
