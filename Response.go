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

	//Send  writes out data to the stream but it is first formatted by the current Encoder.
	Send(string, interface{}) error

	//Error is used to signal to the client that something went wrong.
	//The default action is to send a http status using the first argument. This behaviour can be changed through
	//middleware though.
	Error(int, error) Response

	http.ResponseWriter
}
