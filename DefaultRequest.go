package pelau

import (
	"net/http"
)

type defaultRequest struct {
	request *ModifiedRequest
	dec     map[string]func(Request) Decoder
}

//
func (r *defaultRequest) Params() []string {

	return r.request.params

}

//
func (r *defaultRequest) AddDecoder(typ string, f func(Request) Decoder) {

	r.dec[typ] = f

}

//
func (r *defaultRequest) Retrieve(typ string, i interface{}) error {

	if dec := r.dec[typ]; dec != nil {

		err := dec(r).Decode(i)

		return err

	}

	panic("No Decoder found for type " + typ + "!")

	return nil

}

//
func (r *defaultRequest) Raw() *ModifiedRequest {

	return r.request

}

//DefaultRequest creates a new Request implementation.
func DefaultRequest(req *http.Request) Request {

	return &defaultRequest{&ModifiedRequest{make([]string, 0), req}, make(map[string]func(Request) Decoder)}

}
