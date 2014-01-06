package mid

import (
	"encoding/json"
	"github.com/metasansana/pelau"
)

//JSONInput adds JSON encoding support to the Response object.
func JSONInput(req pelau.Request, res pelau.Response, ctx *pelau.Context) {

	req.AddDecoder(pelau.JSON, func(reader pelau.Reader, i interface{}, f func(error, interface{})) {

		err := json.NewDecoder(reader).Decode(i)
		f(err, i)
	})

	ctx.Next(req, res, ctx)

}
