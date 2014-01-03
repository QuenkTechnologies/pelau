package mid

import (
	"github.com/metasansana/pelau"
	"net/http"
)

//FormParser parse the url query variables and Body of a request.
//If parsing is succesful then the data can be retrieved via pelau.Request.Get()
func FormParser(req pelau.Request, res pelau.Response, ctx *pelau.Context) {

	req.Raw(func(pkgReq *http.Request) {

		err := pkgReq.ParseForm()

		if err != nil {

			println(err.Error())

		}

		ctx.Next(req, res, ctx)

	})

}
