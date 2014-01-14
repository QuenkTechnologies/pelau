package pelau

import (
	"errors"
	"net/http"
)

type defaultResponse struct {
	http.ResponseWriter
}

//Head queues a header up for delivery.
func (r *defaultResponse) Head(key string, value string) Response {

	r.Header().Add(key, value)

	return r

}

//Send writes out an interface to the stream using the interfal formatter is set.
func (r *defaultResponse) Send(mime string, i interface{}) error {

	return errors.New(mime + " encoder not found!")

}

//Redirect is a convenience method for sending locations or redirects.
func (r *defaultResponse) Redirect(url string, status int) Response {

	r.Head("Location", url)

	r.WriteHeader(status)

	return r
}

//Error sends a status header to the client.
func (r *defaultResponse) Error(code int, err error) Response {

	r.WriteHeader(code)

	return r

}

//DefaultResponse creates a Response implementation.
func DefaultResponse(w http.ResponseWriter) Response {

	return &defaultResponse{w}
}
