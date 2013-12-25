package pelau

import (
	"regexp"
)

//RegexRoute provides regex routing support.
type regexRoute struct {
	path   *regexp.Regexp
	action Callback
}

type paramRequest struct {
	Request
}

//Query indicates if we have a regex match or not.
func (s *regexRoute) Query(route string, req Request, res Response) bool {

	if params := s.path.FindStringSubmatch(route); len(params) > 0 {

		s.action(&paramRequest{req}, res)
		return true

	}
	return false
}

//Regex Route constructor.
func Regex(path string, callback Callback) Route {

	return &regexRoute{regexp.MustCompile(path), callback}

}
