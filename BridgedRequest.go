package pelau

import ()

type bridgedRequest struct {
	Request
}

//BridgedRequest constructs a new BridgedRequest
func BridgedRequest(req Request) Request {

	return &bridgedRequest{req}

}
