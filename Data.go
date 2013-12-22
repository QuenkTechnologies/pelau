package pelau

import (
	"net/http"
)

//Data is an interface for formatting data.
type Data interface {
	Output(http.ResponseWriter)
}
