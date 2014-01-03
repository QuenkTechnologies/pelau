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
	onServe func()
	c       *Context
}

//Get adds a GET route
func (r *defaultRouter) Get(route Route) Router {

	r.routes[get] = append(r.routes[get], route)

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
func (r *defaultRouter) Use(m MiddleWare) Router {

	r.c.Add(m)

	return r

}

//ServeHTTP implements the interface from http.Hanlder.
func (r *defaultRouter) ServeHTTP(stdRes http.ResponseWriter, stdReq *http.Request) {

	r.c.Reset()

	r.Use(func(req Request, res Response, c *Context) {

		raw := req.Raw()

		for _, aRoute := range r.routes[raw.Method] {
			if aRoute != nil {
				aRoute.Query(raw.URL.Path, req, res)
			}

		}

	})

	r.c.Next(DefaultRequest(stdReq), DefaultResponse(stdRes), r.c)

}

//Bind binds the server to an interface and starts the app.
func (r *defaultRouter) Bind(addr string, f func()) {

	r.onServe = f

	if f != nil {

		f()

	}

	http.ListenAndServe(addr, r)

}

//DefaultRouter constructs a new Default object.
func DefaultRouter() Router {

	routes := make(routeList)
	routes[get] = make([]Route, 0)
	routes[put] = make([]Route, 0)
	routes[post] = make([]Route, 0)
	routes[del] = make([]Route, 0)
	routes[head] = make([]Route, 0)
	routes[any] = make([]Route, 0)

	return &defaultRouter{routes, func() {}, DefaultContext()}

}
