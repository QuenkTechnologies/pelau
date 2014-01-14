package mime

import (
	"encoding/json"
	"github.com/metasansana/pelau"
)

//JSONResponse adds support for sending a json enconding response.
type JSONResponse struct {
	pelau.Response
}

//Send json encoded data to client if mime is application/json.
func (r *JSONResponse) Send(mime string, i interface{}) error {

	if mime == pelau.JSON {

		r.Head("Content-Type", pelau.JSON)
		return json.NewEncoder(r).Encode(i)
	}
	return r.Response.Send(mime, i)

}

//JSONOutput adds JSON encoding support to the Response object.
func JSONOutput(req pelau.Request, res pelau.Response, ctx *pelau.Context) {

	ctx.Next(req, &JSONResponse{res}, ctx)

}
