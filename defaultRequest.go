package pelau

import (
	"errors"
)

type defaultRequest struct {
	request *ModifiedRequest
}

func (r *defaultRequest) Param(index int) string {

	val := r.request.Params[index]

	return val

}

func (r *defaultRequest) Value(index string) string {

	val := r.request.Form.Get(index)

	return val

}

func (r *defaultRequest) Decode(mime string, i interface{}) error {

	return errors.New(mime + " decoder not found!")

}

func (r *defaultRequest) Raw(f func(*ModifiedRequest)) Request {

	f(r.request)

	return r

}

//DefaultRequest creates a new Request implementation.
func DefaultRequest(req *ModifiedRequest) Request {

	return &defaultRequest{req}

}
