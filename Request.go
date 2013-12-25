package pelau

import (
	"net/http"
)

//Request is the inteface for incomming http infomation.
type Request interface {
	Params() []string

	Raw() *http.Request
}
