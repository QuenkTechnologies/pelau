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

//Write writes out data to the client
func (r *defaultResponse) Write(d Data) {

	d.Output(r.w)

}

//Created sends the created http status.
func (r *defaultResponse) Created(loc string) {

	//if loc != nil {

	//	r.Redirect(loc, 304)

	//	} else {

	//		r.Status(304)
	//		r.w.Write(nil)
	//	}

}

//Conflict sends conflict status
func (r *defaultResponse) Conflict() {

	r.Status(409)
	r.w.Write(nil)

}

//NotFound sends a 404 status
func (r *defaultResponse) NotFound() {

	r.Status(404)
	r.w.Write(nil)

}

//Ok sends the 200 status
func (r *defaultResponse) Ok() {

	r.Status(200)
	r.w.Write(nil)

}

//DefaultResponse creates a Response implementation.
func DefaultResponse(w http.ResponseWriter) Response {

	return &defaultResponse{w}

}
