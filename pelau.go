/**

Pelau

Pelau is a small framework for making http REST api servers in Go. It adopts some concepts from frameworks like
Express, Connect, Sinatra and friends. It is not meant to be the end all framework so it probably won't do everything
you want. However if all you want is to respond to some http request you've got the right framework :).

*/

package pelau

import (
	"io"
	"net/http"
)

/*
  This file contains some types that are too small or annoying to put in their own file.
*/

//ModifiedRequest is a wrapper around http.Request that includes the result of a patten match on the path part of the request uri.
type ModifiedRequest struct {
	Params []string
	*http.Request
}

//Reader is here for convenience
type Reader interface {
	io.Reader
}

//Writer is here for convenience
type Writer interface {
	io.Writer
}

//Callback coresponds to Callback's signature
type Callback func(req Request, res Response)

//Decoder funcs are used to decode raw incomming request data.
type Decoder func(Reader, interface{}, func(error, interface{}))

//Encoder funcs are used to encode raw outgoing data.
type Encoder func(Writer, interface{}, func(error, int))
