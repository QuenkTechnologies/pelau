package pelau_test

/*
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
		get     pelau.Route
		post    pelau.Route
		mockCtl *gomock.Controller
		f       func()
	)

	BeforeEach(func() {

		router = pelau.DefaultRouter()
		mockCtl = gomock.NewController(gomocktestreporter.New())
		server = httptest.NewServer(router)
		route = pelau.NewMockRoute(mockCtl)
		get = pelau.NewMockRoute(mockCtl)
		post = pelau.NewMockRoute(mockCtl)
		count = 0
		f = func() {

			router.Get(get)
			router.Post(post)
			router.Put(route)
			router.Delete(route)
			router.Head(route)

		}

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

	Context("when routing", func() {

		Describe("on GET requests", func() {

			It("execute the route if the request is a GET one", func() {

				f()

				http.Get(server.URL)

			})
		})
	})
})*/
