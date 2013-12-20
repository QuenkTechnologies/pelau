package pelau

import (
	"net/http"
)

const (
	get  = "GET"
	put  = "PUT"
	post = "POST"
	del  = "DELETE"
	head = "HEAD"
)

type list map[string][]Route

//Server is the object that sets up routing.
type Server struct {
	routes list
}

//Regex is a factory method for regex route support.
func (r *Server) Regex(regex string, callback Callback) Route {

	return NewRegexRoute(regex, callback)
}

//Static is a factory method for static route support.
func (r *Server) Static(path string, callback Callback) Route {

	return NewStaticRoute(path, callback)
}

//Get adds a GET route
func (r *Server) Get(route Route) *Server {

	r.routes[get] = append(r.routes["GET"], route)

	return r

}

//Put adds a PUT route
func (r *Server) Put(route Route) *Server {

	r.routes[put] = append(r.routes[put], route)

	return r

}

//Post adds a POST route
func (r *Server) Post(route Route) *Server {

	r.routes[post] = append(r.routes[post], route)

	return r

}

//Delete adds a DELETE route.
func (r *Server) Delete(route Route) *Server {

	r.routes[del] = append(r.routes[del], route)

	return r

}

//Head adds a HEAD route.
func (r *Server) Head(route Route) *Server {

	r.routes[head] = append(r.routes[head], route)

	return r

}

//ServeHTTP implements the interface from http.Hanlder.
func (r *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	for _, aRoute := range r.routes[req.Method] {

		if result := aRoute.Query(req.URL.Path); result {

		}

	}

}

//Bind binds the server to an interface and starts the app.
func (r *Server) Bind(addr string) {

	http.ListenAndServe(addr, r)

}

//New constructs a new Server object.
func New() *Server {

	routes := make(list)
	routes["GET"] = make([]Route, 0)
	routes["PUT"] = make([]Route, 0)
	routes["POST"] = make([]Route, 0)
	routes["DELETE"] = make([]Route, 0)
	routes["HEAD"] = make([]Route, 0)

	return &Server{routes}

}
