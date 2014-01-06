package pelau

import (
	"errors"
	"net/http"
)

type defaultResponse struct {
	http.ResponseWriter
	enc map[string]Encoder
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

//Send writes out an interface to the stream using the interfal formatter is set.
func (r *defaultResponse) Send(mime string, i interface{}, f func(error, int)) Response {

	if enc := r.enc[mime]; enc == nil {

		f(errors.New("No Encoder found for type "+mime+"!"), 0)

	} else {

		enc(r, i, f)

	}

	return r

}

//Redirect is a convenience method for sending locations or redirects.
func (r *defaultResponse) Redirect(url string, status int) Response {

	r.Head("Location", url)

	r.WriteHeader(status)

	return r
}

//DefaultResponse creates a Response implementation.
func DefaultResponse(w http.ResponseWriter) Response {

	return &defaultResponse{w, make(map[string]Encoder)}
}
