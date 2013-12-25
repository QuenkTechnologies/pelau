package pelau

//Encoder represents an interface for encoding data in some specific format.
type Encoder interface {
	Encode(interface{}) error
}
