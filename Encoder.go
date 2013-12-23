package pelau

type Encoder interface {
	Encode(interface{}) error
}
