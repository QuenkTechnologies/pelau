package pelau

//Request is the interface used to retrieve information about inbound http requests.
type Request interface {

	//Param provides the strings matched (if any) after a regex match on the uri path.
	Param(int) string

	//Value retrieves the value of a query variable or a form body key.
	//In order for this to work you must include the BodyParser middleware.
	Value(string) string

	//Decode parses the contents of the request body using preconfigured middleware.
	Decode(mime string, i interface{}) error

	//Raw accepts a callback that is given access to the ModifiedRequest struct used internally.
	Raw(func(*ModifiedRequest)) Request
}
