package pelau

//Request is the interface used to retrieve information about inbound http requests.
type Request interface {

	//Param provides the strings matched (if any) after a regex match on the uri path.
	Param(int) string

	//Get retrieves the value of a query variable or a form body key.
	//In order for this to work you must include the BodyParser middleware.
	Get(string) string

	//Retrieve attempts to parse the contents of the request body but in a specified format.
	//Calls to Retrieve() passes the suplied interface and callback  to the associated Decoder (see AddDecoder()).
	Retrieve(mime string, i interface{}, callback func(error, interface{}))

	//AddDecoder registers a func that will be used by Read() to decode the request body.
	AddDecoder(string, Decoder) Request

	//Raw accepts a callback that is given access to the ModifiedRequest struct used internally.
	Raw(func(*ModifiedRequest)) Request
}
