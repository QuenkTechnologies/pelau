package pelau

//Response is the interface for manipulating outgoing http information.
type Response interface {
	Status(int) Response

	Write(Data)

	Created(string)

	Conflict()

	NotFound()

	Ok()

	//	Redirect(string, int)

	//	Replace(Response)
}
