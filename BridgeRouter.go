package pelau

//BridgeRouter is used to create a bridge pattern between the framework and customised routers.
type BridgeRouter struct {
	router Router
}

func (b BridgeRouter) Use(c Callback) {

	b.router.Use(c)

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

//Init is the first step to using the framework. It expects a Router implementation to use.
func Init(r Router) Router {

	return &BridgeRouter{r}

}
