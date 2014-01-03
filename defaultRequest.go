package pelau

import (
	"errors"
	"net/http"
)

type defaultRequest struct {
	request  *ModifiedRequest
	decoders map[string]func(Request) Decoder
}

func (r *defaultRequest) Param(index int) string {

	val := r.request.params[index]

	return val

}

func (r *defaultRequest) Get(index string) string {

	val := r.request.Form.Get(index)

	return val

}

func (r *defaultRequest) ParseBody(mime string, i interface{}, f func(error, interface{})) Request {

	var err error
	if decoder := r.decoders[mime]; decoder != nil {

		err = decoder(r).Decode(i)

	} else {

		err = errors.New("No decoder found for type " + mime + ".")

	}

	f(err, i)

	return r

}

func (r *defaultRequest) Register(mime string, f func(Request) Decoder) Request {

	r.decoders[mime] = f
	return r

}

func (r *defaultRequest) Raw(f func(*ModifiedRequest)) Request {

	f(r.request)

	return r

}

//DefaultRequest creates a new Request implementation.
func DefaultRequest(req *http.Request) Request {

	return &defaultRequest{&ModifiedRequest{make([]string, 0), req}, make(map[string]func(Request) Decoder)}

}
