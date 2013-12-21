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
	any  = "ANY"
)

type routeList map[string][]Route

//DefaultRouter is the object that sets up routing.
type defaultRouter struct {
	routes  routeList
	onServe BindFunc
}

//Get adds a GET route
func (r *defaultRouter) Get(route Route) Router {

	r.routes[get] = append(r.routes["GET"], route)

	return r

}

//Put adds a PUT route
func (r *defaultRouter) Put(route Route) Router {

	r.routes[put] = append(r.routes[put], route)

	return r

}

//Post adds a POST route
func (r *defaultRouter) Post(route Route) Router {

	r.routes[post] = append(r.routes[post], route)

	return r

}

//Delete adds a DELETE route.
func (r *defaultRouter) Delete(route Route) Router {

	r.routes[del] = append(r.routes[del], route)

	return r

}

//Head adds a HEAD route.
func (r *defaultRouter) Head(route Route) Router {

	r.routes[head] = append(r.routes[head], route)

	return r

}

//Use allows middleware to be added to this router.
func (r *defaultRouter) Use(route Route) Router {

	r.routes[any] = append(r.routes[any], route)

	return r

}

//ServeHTTP implements the interface from http.Hanlder.
func (r *defaultRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	for _, aUse := range r.routes[any] {

		aUse.Query(req.URL.Path, res, req)

	}

	for _, aRoute := range r.routes[req.Method] {

		aRoute.Query(req.URL.Path, res, req)

	}

}

//Bind binds the server to an interface and starts the app.
func (r *defaultRouter) Bind(addr string, f BindFunc) {

	r.onServe = f

	http.ListenAndServe(addr, r)

	if f != nil {

		f()

	}

}

//DefaultRouter constructs a new Server object.
func DefaultRouter() Router {

	routes := make(routeList)
	routes[get] = make([]Route, 0)
	routes[put] = make([]Route, 0)
	routes[post] = make([]Route, 0)
	routes[del] = make([]Route, 0)
	routes[head] = make([]Route, 0)
	routes[any] = make([]Route, 0)

	return &defaultRouter{routes, DefaultBindFunc()}

}
