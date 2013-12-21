package pelau

import (
	"net/http"
)

//Router interface defines methods that are used for setting up a pelau app and running it.
type Router interface {
	Use(r Route) Router

	Get(r Route) Router

	Post(r Route) Router

	Put(r Route) Router

	Delete(r Route) Router

	Head(r Route) Router

	Bind(string, BindFunc)

	ServeHTTP(http.ResponseWriter, *http.Request)
}
