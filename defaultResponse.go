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

//UseEncoder sets the Encoding
func (r *defaultResponse) AddEncoder(typ string, enc Encoder) Response {

	r.enc[typ] = enc

	return r

}

//Stream writes out an interface to the stream using the interfal formatter is set.
func (r *defaultResponse) Stream(mime string, i interface{}) (error, int) {

	return error.New(mime + " encoder not found!"), 0

}

//Redirect is a convenience method for sending locations or redirects.
func (r *defaultResponse) Redirect(url string, status int) Response {

	r.Head("Location", url)

	r.WriteHeader(status)

	return r
}

//DefaultResponse creates a Response implementation.
func DefaultResponse(w http.ResponseWriter) Response {

	return &defaultResponse{w}
}
