package pelau

import (
	"net/http"
)

//Response is the interface for manipulating outgoing http information.
type Response interface {

	//Head sends an HTTP header to the client.
	Head(string, string) Response

	//Redirect sends a location header to the client
	Redirect(string, int) Response

	//Stream  writes out data to the stream but it is first formatted by the current Encoder.
	Stream(string, interface{}) (error, int)

	http.ResponseWriter
}
