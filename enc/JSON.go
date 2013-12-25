package mid

import (
	"encoding/json"
	"github/metasansana/pelau"
)

//JSONEncoding adds JSON encoding support to the Response object.
func JSONEncoding(req pelau.Request, res pelau.Response, ctx *pelau.Content) {

	res.AddEncoder(func(res pelau.Response) Encoder {

		res.Head("Content-Type", "application/json")
		return json.NewEncoder(res)

	})
	ctx.Next(req, res, ctx)

}
