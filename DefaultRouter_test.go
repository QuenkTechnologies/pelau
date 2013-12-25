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

var _ = Describe("DefaultRouter", func() {

	var (
		router  pelau.Router
		server  *httptest.Server
		count   int
		route   pelau.Route
		mockCtl *gomock.Controller
	)

	BeforeEach(func() {

		router = pelau.DefaultRouter()
		mockCtl = gomock.NewController(gomocktestreporter.New())
		server = httptest.NewServer(router)
		route = pelau.NewMockRoute(mockCtl)
		count = 0

	})

	AfterEach(func() {

		defer server.Close()

	})

	Describe("when there is middleware", func() {

		It("should activate any middleware for all http calls", func() {
			router.Use(func(req pelau.Request, res pelau.Response, c *pelau.Context) {

				count++

			})

			http.Get(server.URL)
			http.Head(server.URL)
			Expect(count).To(Equal(2))

		})

	})

	Describe("when routing", func() {

		It("should work on  different request types", func() {

			router.Get(route)
			router.Head(route)

			http.Get(server.URL)
			http.Head(server.URL)

		})
	})
})
