package pelau

import ()

//BridgeRouter is used to create a bridge pattern between the framework and customised routers.
type BridgeRouter struct {
	Router
}

//Cook is the first step to using the framework. It expects a Router implementation to use.
func Cook(r Router) Router {

	return &BridgeRouter{r}

}
