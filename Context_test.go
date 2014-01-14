package pelau_test

import (
	"code.google.com/p/gomock/gomock"
	"github.com/quenktech/pelau"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/thirdparty/gomocktestreporter"
	. "github.com/onsi/gomega"
)

var _ = Describe("Context", func() {

	var (
		req     pelau.Request
		res     pelau.Response
		ctx     *pelau.Context
		mockCtl *gomock.Controller
	)

	BeforeEach(func() {
		mockCtl = gomock.NewController(gomocktestreporter.New())
		req = pelau.NewMockRequest(mockCtl)
		res = pelau.NewMockResponse(mockCtl)
		ctx = pelau.DefaultContext()

	})

	It("should call the middleware functions", func() {

		var a, b, c string

		ctx.Add(func(req pelau.Request, res pelau.Response, cx *pelau.Context) {

			a = "A"

			cx.Next(req, res, cx)

		})
		ctx.Add(func(req pelau.Request, res pelau.Response, cx *pelau.Context) {

			b = "B"
			cx.Next(req, res, cx)

		})
		ctx.Add(func(req pelau.Request, res pelau.Response, cx *pelau.Context) {

			c = "C"
			cx.Next(req, res, cx)

		})

		ctx.Next(req, res, ctx)

		Expect(a + b + c).Should(Equal("ABC"))

	})

})
