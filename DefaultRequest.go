package pelau

import (
	"net/http"
	"net/url"
)

type defaultRequest struct {
	request *http.Request
	params  []string
}

//Method provides the method value of the request.
func (r defaultRequest) Method() string {

	return r.request.Method

}

//Url returns the url type for this request.
func (r defaultRequest) Url() *url.URL {

	return r.request.URL

}

//
func (r defaultRequest) Params() []string {

	return r.params

}

//
func (r defaultRequest) Raw() *http.Request {

	return r.request

}

//DefaultRequest creates a new Request implementation.
func DefaultRequest(req *http.Request, params []string) Request {

	return &defaultRequest{req, params}

}
