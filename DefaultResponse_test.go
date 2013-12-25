package pelau_test

import (
	"github.com/metasansana/pelau"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("DefaultResponse", func() {

	var (
		server *httptest.Server
		f      func(pelau.Response)
	)

	BeforeEach(func() {

		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			f(pelau.DefaultResponse(w))

		}))

	})

	AfterEach(func() {

		defer server.Close()

	})

	Describe("when sending headers", func() {
		It("should allow users to send headers to clients", func() {
			poweredby := "pelau"
			itest := "1"
			allow := "GET, HEAD, PUT, DELETE, POST"
			f = func(res pelau.Response) {

				res.Head("X-Powered-By", poweredby).
					Head("X-Integer-Test", itest).
					Head("Allow", allow)

			}

			result, _ := http.Get(server.URL)
			h := result.Header
			Expect(h.Get("X-Powered-By")).To(Equal(poweredby))
			Expect(h.Get("X-Integer-Test")).To(Equal(itest))
			Expect(h.Get("Allow")).To(Equal(allow))

		})
	})
})
