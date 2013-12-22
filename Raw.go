package pelau

import (
	"net/http"
)

//
type raw struct {
	bits []byte
}

//
func (self raw) Output(w http.ResponseWriter) {

	w.Write(self.bits)
}

func Raw(bits []byte) Data {

	return &raw{bits}

}
