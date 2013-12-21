package pelau

import (
	"net/http"
)

//BridgeRouter is used to create a bridge pattern between the framework and customised routers.
type BridgeRouter struct {
	router Router
}

func (b BridgeRouter) Use(r Route) Router {

	b.router.Use(r)

	return b

}

func (b BridgeRouter) Get(r Route) Router {

	b.router.Get(r)
	return b

}

func (b BridgeRouter) Post(r Route) Router {

	b.router.Post(r)
	return b
}

func (b BridgeRouter) Put(r Route) Router {

	b.router.Put(r)
	return b

}

func (b BridgeRouter) Delete(r Route) Router {

	b.router.Delete(r)
	return b

}

func (b BridgeRouter) Head(r Route) Router {

	b.router.Head(r)
	return b

}

func (b BridgeRouter) Bind(addr string) {

	b.router.Bind(addr)

}

func (b BridgeRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	b.router.ServeHTTP(w, r)

}

//Cook is the first step to using the framework. It expects a Router implementation to use.
func Cook(r Router) Router {

	return &BridgeRouter{r}

}
