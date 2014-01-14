package mime

import (
	"encoding/json"
	"github.com/metasansana/pelau"
)

//JSONRequest adds support for reading json requests.
type JSONRequest struct {
	pelau.Request
}

//Decode overwrites pelau.Request.Decode()
func (r *JSONRequest) Decode(mime string, i interface{}) error {

	if mime == pelau.JSON {

		var decoder *json.Decoder

		r.Raw(func(mod *pelau.ModifiedRequest) {

			decoder = json.NewDecoder(mod.Body)

		})

		return decoder.Decode(i)

	}

	return r.Request.Decode(mime, i)

}

//JSONInput adds JSON encoding support to the Response object.
func JSONInput(req pelau.Request, res pelau.Response, ctx *pelau.Context) {

	ctx.Next(&JSONRequest{req}, res, ctx)

}
