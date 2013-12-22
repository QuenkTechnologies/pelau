package pelau

import (
	"net/http"
)

//Router interface defines methods that are used for setting up a pelau app and running it.
type Router interface {
	Use(MiddleWare) Router

	Get(Route) Router

	Post(Route) Router

	Put(Route) Router

	Delete(Route) Router

	Head(Route) Router

	Bind(string, BindFunc)

	ServeHTTP(http.ResponseWriter, *http.Request)
}
