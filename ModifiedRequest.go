package pelau

import (
	"net/http"
)

//ModifiedRequest is a wrapper around http.Request that includes pattern match results on the url.
type ModifiedRequest struct {
	params []string
	*http.Request
}
