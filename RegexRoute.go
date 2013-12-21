package pelau

import (
	"net/http"
	"regexp"
)

//RegexRoute provides regex routing support.
type regexRoute struct {
	path   *regexp.Regexp
	action Callback
}

//Query indicates if we have a regex match or not.
func (s *regexRoute) Query(route string, w http.ResponseWriter, r *http.Request) bool {

	if params := s.path.FindStringSubmatch(route); len(params) > 0 {

		s.action(DefaultRequest(r, params), DefaultResponse(w))

		return true

	}
	return false
}

//Regex Route constructor.
func Regex(path string, callback Callback) Route {

	return &regexRoute{regexp.MustCompile(path), callback}

}
