package pelau

//Request is the interface used to retrieve information about inbound http requests.
type Request interface {

	//Param provides the strings matched (if any) after a regex match on the uri path.
	Param(int) string

	//Get retrieves the value of a query variable or a form body key.
	//In order for this to work you must include the BodyParser middleware.
	Get(string) string

	//Read attempts to parse the contents of the request body but in a specified format.
	//Calls to Read() passes the suplied interface to the associated Decoder (see Register()) and invokes a supplied Callback
	//after decoding.
	Read(string, interface{}, func(error, interface{})) Request

	//Register registers a func that will be used by Read() to decode the request body.
	Register(string, func(Request) Decoder) Request

	//Raw accepts a callback that is given access to the ModifiedRequest struct used internally.
	Raw(func(*ModifiedRequest)) Request
}
