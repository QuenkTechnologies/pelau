package pelau

import (
	"net/http"
)

type defaultResponse struct {
	http.ResponseWriter
}

//DefaultResponse creates a Response implementation.
func DefaultResponse(w http.ResponseWriter) Response {

	return &defaultResponse{w}

}
