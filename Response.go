package pelau

//Response is the interface for manipulating outgoing http information.
type Response interface {
	Status(int) Response

	Header(string, string) Response

	SetEncoder(Encoder) Response

	WriteData(interface{}) Response

	Write([]byte) (int, error)

	Redirect(string, int) Response
}
