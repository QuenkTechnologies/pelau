package pelau_test

import (
	"github.com/metasansana/pelau"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"net/url"
)

var _ = Describe("DefaultRequest", func() {

	var (
		server *httptest.Server
		f      func(pelau.Request)
		init   func(http.ResponseWriter, *http.Request)
	)
	initWithFormParser := func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

	}

	BeforeEach(func() {

		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if init != nil {
				init(w, r)
			}

			f(pelau.DefaultRequest(&pelau.ModifiedRequest{[]string{}, r}))

		}))

	})

	AfterEach(func() {

		f = nil
		defer server.Close()

	})

	Describe("request variables", func() {

		BeforeEach(func() {

			init = initWithFormParser
		})

		Context("when a Get request", func() {

			It("should be available via Get()", func() {

				f = func(req pelau.Request) {

					Expect(req.Get("first")).To(Equal("1"))

				}

				http.Get(server.URL + "?first=1&second=two")
			})
		})
		Context("when a Post request", func() {

			It("should be available via Get()", func() {

				f = func(req pelau.Request) {

					Expect(req.Get("first")).To(Equal("1"))

				}
				http.PostForm(server.URL, url.Values{"first": {"1"}})

			})
		})

	})

	Describe("request uri params", func() {

		var req pelau.Request
		BeforeEach(func() {

			mreq := pelau.ModifiedRequest{}

			mreq.Params = []string{"wholeurl", "m1", "m2"}

			req = pelau.DefaultRequest(&mreq)
		})

		It("should provide matches from the regex applied to the url", func() {

			Expect(req.Param(0)).To(Equal("wholeurl"))
			Expect(req.Param(2)).To(Equal("m2"))
		})

	})

	Describe("calling the Raw() func", func() {

		It("should provide access to the ModifiedRequest", func() {

			f = func(req pelau.Request) {

				req.Raw(func(m *pelau.ModifiedRequest) {

					Expect(m).ShouldNot(BeNil())

				})

			}

			http.Get(server.URL)

		})

	})

	Describe("using decoding features", func() {

		It("should allow users to register a decoder func", func() {

			f = func(req pelau.Request) {

				req.AddDecoder("text/plain", func(reader pelau.Reader, i interface{}, cb func(error, interface{})) {

					maap := i.(map[string]string)

					maap["test"] = "TEST"

					cb(nil, maap)

				})

				req.Retrieve("text/plain", make(map[string]string), func(err error, i interface{}) {

					v := i.(map[string]string)

					Expect(v["test"]).To(Equal("TEST"))

				})

			}

			http.Get(server.URL)

		})

	})

})
