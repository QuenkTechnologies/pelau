package pelau

import (
	"net/url"
)

//Request is the inteface for incomming http infomation.
type Request interface {
	Method() string

	Url() *url.URL

	Params() []string
}
