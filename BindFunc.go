package pelau

//BindFunc is the type of the callback used when we start listening for requests.
type BindFunc func()

//DefaultBindFunc provides a default implementation of BindFunc when a concrete implementation is not desired.
func DefaultBindFunc() BindFunc {

	return func() {}

}
