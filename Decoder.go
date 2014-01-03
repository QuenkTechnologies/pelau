package pelau

import ()

type Decoder interface {
	Decode(interface{}) error
}
