package pelau

import ()

//Context represents the execution context when serving http handlers
type Context struct {
	mid []MiddleWare
	ptr int
}

//DefaultContext is a Context costructor
func DefaultContext() *Context {

	return &Context{make([]MiddleWare, 0), 0}

}

//Add includes middleware into the execution context.
func (ctx *Context) Add(m MiddleWare) {

	ctx.mid = append(ctx.mid, m)
}

//Reset resets the internal pointer over the list of middleware.
func (ctx *Context) Reset() {

	ctx.ptr = 0

}

//Next executes the next middleware in the list (if any).
func (ctx *Context) Next(req Request, res Response, c *Context) {
	if ctx.ptr <= len(ctx.mid)-1 {
		f := ctx.mid[ctx.ptr]

		ctx.ptr++

		f(req, res, c)

	}

}
