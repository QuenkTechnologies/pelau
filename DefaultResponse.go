package pelau

import (
	"net/http"
)

type defaultResponse struct {
	w http.ResponseWriter
}

//Status sets the status of the response
func (r *defaultResponse) Status(code int) Response {

	r.w.WriteHeader(code)

	return r

}

//Header queues a header up for delivery.
func (r *defaultResponse) Header(key string, value string) Response {

	r.w.Header().Add(key, value)

	return r

}

//SetEncoder sets the Encoding
func (r *defaultResponse) SetEncoder(e Encoder) Response {

	return r

}

//WriteData writes out an interface to the stream using the interfal formatter is set.
func (r *defaultResponse) WriteData(interface{}) Response {

	return r

}

//Write writes out data to the client
func (r *defaultResponse) Write(data []byte) (int, error) {

	return r.w.Write(data)

}

//Redirect is a convenience method for sending locations or redirects.
func (r *defaultResponse) Redirect(url string, status int) Response {

	return r
}

//DefaultResponse creates a Response implementation.
func DefaultResponse(w http.ResponseWriter) Response {

	return &defaultResponse{w}

}
