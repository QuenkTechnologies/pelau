package mid

import (
	"encoding/json"
	"github.com/metasansana/pelau"
)

//JSONInput adds JSON encoding support to the Response object.
func JSONInput(req pelau.Request, res pelau.Response, ctx *pelau.Context) {

	req.AddDecoder(pelau.JSON, func(current pelau.Request) pelau.Decoder {

		return json.NewDecoder(current.Raw().Body)

	})

	ctx.Next(req, res, ctx)

}
