package pelau

//Router interface defines methods that are used for setting up a pelau app and running it.
type Router interface {
	Use(c Callback)

	Get(r Route) Router

	Post(r Route) Router

	Put(r Route) Router

	Delete(r Route) Router

	Head(r Route) Router

	Bind(addr string)
}
