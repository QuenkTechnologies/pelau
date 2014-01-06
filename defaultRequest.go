package pelau

import (
	"errors"
)

type defaultRequest struct {
	request  *ModifiedRequest
	decoders map[string]Decoder
}

func (r *defaultRequest) Param(index int) string {

	val := r.request.Params[index]

	return val

}

func (r *defaultRequest) Get(index string) string {

	val := r.request.Form.Get(index)

	return val

}

func (r *defaultRequest) Retrieve(mime string, i interface{}, f func(error, interface{})) {

	if decode := r.decoders[mime]; decode == nil {

		f(errors.New("No decoder found for type "+mime+"!"), i)

	} else {

		r.Raw(func(m *ModifiedRequest) {

			decode(m.Body, i, f)

		})

	}

}

func (r *defaultRequest) AddDecoder(mime string, f Decoder) Request {

	r.decoders[mime] = f
	return r

}

func (r *defaultRequest) Raw(f func(*ModifiedRequest)) Request {

	f(r.request)

	return r

}

//DefaultRequest creates a new Request implementation.
func DefaultRequest(req *ModifiedRequest) Request {

	return &defaultRequest{req, make(map[string]Decoder)}

}
