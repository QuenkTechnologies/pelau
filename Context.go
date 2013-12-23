package pelau

type Context struct {
	mid  []MiddleWare
	ptr  int
	ptr2 int
}

func DefaultContext() *Context {

	return &Context{make([]MiddleWare, 0), 0, 0}

}

//
func (self *Context) Add(m MiddleWare) {

	self.mid = append(self.mid, m)
	self.ptr++
}

func (self *Context) Reset() {

	self.ptr = 0
	self.ptr2 = 0

}

//
func (self *Context) Next(req Request, res Response, c *Context) {

	var f MiddleWare

	if self.ptr2 < self.ptr {

		f = self.mid[self.ptr2]

		if f != nil {

			f(req, res, c)

		}

	}
	self.ptr2++

}
