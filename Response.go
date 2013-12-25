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

	//AddEncoder sets the Encoder that will be used by future calls to WriteData
	AddEncoder(string, func(Response) Encoder) Response

	//Send  writes out data to the stream but it is first formatted by the current Encoder.
	Send(string, interface{}) Response

	http.ResponseWriter
}
