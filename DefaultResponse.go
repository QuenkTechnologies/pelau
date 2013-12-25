package pelau

import (
	"net/http"
)

type defaultResponse struct {
	http.ResponseWriter
	enc map[string]func(Response) Encoder
}

//Head queues a header up for delivery.
func (r *defaultResponse) Head(key string, value string) Response {

	r.Header().Add(key, value)

	return r

}

//UseEncoder sets the Encoding
func (r *defaultResponse) AddEncoder(typ string, enc func(Response) Encoder) Response {

	r.enc[typ] = enc

	return r

}

//Send writes out an interface to the stream using the interfal formatter is set.
func (r *defaultResponse) Send(typ string, i interface{}) Response {

	if enc := r.enc[typ]; enc != nil {

		err := enc(r).Encode(i)

		if err != nil {

			println(err)

		}

	} else {

		panic("No Encoder found for type " + typ + "!")

	}

	return r

}

//Redirect is a convenience method for sending locations or redirects.
func (r *defaultResponse) Redirect(url string, status int) Response {

	r.WriteHeader(status)

	r.Head("Location", url)

	return r
}

//DefaultResponse creates a Response implementation.
func DefaultResponse(w http.ResponseWriter) Response {

	return &defaultResponse{w, make(map[string]func(Response) Encoder)}
}
