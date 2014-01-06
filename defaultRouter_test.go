package pelau_test

import (
	"code.google.com/p/gomock/gomock"
	"github.com/metasansana/pelau"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/thirdparty/gomocktestreporter"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

type rt struct {
	route string
	flag  bool
}

func (r *rt) Query(path string, req pelau.Request, res pelau.Response) bool {

	r.flag = true
	return true

}

var _ = Describe("DefaultRouter", func() {
	var (
		router  pelau.Router
		server  *httptest.Server
		mockCtl *gomock.Controller
		called  bool
	)

	BeforeEach(func() {

		router = pelau.DefaultRouter()
		mockCtl = gomock.NewController(gomocktestreporter.New())
		server = httptest.NewServer(router)
		called = false

	})

	AfterEach(func() {

		defer server.Close()

	})

	Describe("using middleware", func() {

		BeforeEach(func() {
			server = httptest.NewServer(router)
			router.Use(func(req pelau.Request, res pelau.Response, c *pelau.Context) {

				called = true

			})
			router.Use(func(req pelau.Request, res pelau.Response, c *pelau.Context) {

				called = false

			})
			router.Use(func(req pelau.Request, res pelau.Response, c *pelau.Context) {

				called = true

			})

		})
		test := func(f func()) {

			It("should activate middleware", func() {
				f()
				Expect(called).To(Equal(true))
			})

		}
		Context("with GET requests", func() {

			test(func() {
				http.Get(server.URL)
			})

		})
		Context("with POST requests", func() {

			test(func() {
				http.PostForm(server.URL, nil)
			})

		})
		Context("with HEAD requests", func() {

			test(func() {
				http.Head(server.URL)
			})

		})

	})

	Describe("using routing", func() {

		It("should trigger matching GET routes", func() {
			r := &rt{"/test", false}
			router.Get(r)
			http.Get(server.URL + "/test")
			Expect(r.flag).To(Equal(true))

		})
		It("should trigger matching POST routes", func() {
			r := &rt{"/test", false}
			router.Post(r)
			http.PostForm(server.URL+"/test", nil)
			Expect(r.flag).To(Equal(true))

		})
	})
})
