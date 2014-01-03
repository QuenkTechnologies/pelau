package pelau

import ()

//Request is the inteface for incomming http infomation.
type Request interface {
	Params() []string

	AddDecoder(string, func(Request) Decoder)

	Retrieve(string, interface{}) error

	Raw() *ModifiedRequest
}
