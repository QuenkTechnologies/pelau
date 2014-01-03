package pelau

//Request is the interface used to retrieve information about inbound http requests.
type Request interface {

	//Param returns the string indexed (if any) after a succesful match of a Regex route.
	Param(int) string

	//Get retrieves the value of a url query variable or a form key.
	//In order for this to work you must include the FormParser middleware.
	Get(string) string

	//ParseBody attempts to parse the contents of the request body and as the type specified by the first argument.
	//After passing the body into the given interface a callback is called.
	ParseBody(string, interface{}, func(error, interface{})) Request

	//Register registers a func that will be used when calls to Read() match the first argument.
	Register(string, func(Request) Decoder) Request

	//Raw provides access to the raw http.Request type.
	Raw(func(*ModifiedRequest)) Request
}
