package pelau

import (
	"net/http"
)

type defaultRequest struct {
	request *http.Request
	params  []string
}

//
func (r *defaultRequest) Params() []string {

	return r.params

}

//
func (r *defaultRequest) Raw() *http.Request {

	return r.request

}

//DefaultRequest creates a new Request implementation.
func DefaultRequest(req *http.Request) Request {

	return &defaultRequest{req, make([]string, 0)}

}
