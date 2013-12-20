package pelau

import (
	"regexp"
)

//RegexRoute provides regex routing support.
type RegexRoute struct {
	path   *regexp.Regexp
	action Callback
}

//Query indicates if we have a regex match or not.
func (s *RegexRoute) Query(route string) bool {

	if params := s.path.FindStringSubmatch(route); len(params) > 0 {

		return true

	}
	return false
}

//Execute executes this route.
func (s *RegexRoute) Execute(req *Request, res *Response) {

	s.action(req, res)

}

//NewRegexRoute constructor.
func NewRegexRoute(path string, callback Callback) Route {

	return &RegexRoute{regexp.MustCompile(path), callback}

}
