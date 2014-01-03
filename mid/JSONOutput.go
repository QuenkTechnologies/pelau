package mid

import (
	"encoding/json"
	"github.com/metasansana/pelau"
)

//JSONOutput adds JSON encoding support to the Response object.
func JSONOutput(req pelau.Request, res pelau.Response, ctx *pelau.Context) {

	res.AddEncoder(pelau.JSON, func(res pelau.Response) pelau.Encoder {

		res.Head("Content-Type", pelau.JSON)
		return json.NewEncoder(res)

	})
	ctx.Next(req, res, ctx)

}
