package pelau

import ()

type bridgedResponse struct {
	Response
}

//BridgedResponse constructs a new Response
func BridgedResponse(res Response) Response {

	return &bridgedResponse{res}

}
