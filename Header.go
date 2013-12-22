package pelau

import (
	"net/http"
)

type header struct {
	key   string
	value string
}

//Output
func (h header) Output(w http.ResponseWriter) {

	w.Header().Add(h.key, h.value)

}

//Header constructs a new Header Data
func Header(key, value string) Data {

	return &header{key, value}

}
