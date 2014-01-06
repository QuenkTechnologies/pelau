package mid

import (
	"encoding/json"
	"github.com/metasansana/pelau"
)

//JSONOutput adds JSON encoding support to the Response object.
func JSONOutput(req pelau.Request, res pelau.Response, ctx *pelau.Context) {

	res.AddEncoder(pelau.JSON, func(w pelau.Writer, i interface{}, f func(error, int)) {

		res.Head("Content-Type", pelau.JSON)

		bits, err := json.NewEncoder(w).Encode(i)

		f(err, bits)

	})
	ctx.Next(req, res, ctx)

}
