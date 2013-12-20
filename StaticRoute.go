package pelau

import ()

//StaticRoute provides static routing support.
type StaticRoute struct {
	path   string
	action Callback
}

//Query looks up whether the current http path matches the static route..
func (s *StaticRoute) Query(route string) bool {

	if s.path == route {

		return true

	}

	return false

}

//Execute executes this route.
func (s *StaticRoute) Execute(req *Request, res *Response) {

	s.action(req, res)

}

//NewStaticRoute constructor.
func NewStaticRoute(path string, callback Callback) Route {

	return &StaticRoute{path, callback}

}
