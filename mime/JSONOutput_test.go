package mime_test

import (
	"encoding/json"
	"github.com/quenktech/pelau"
	"github.com/quenktech/pelau/mime"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("JSONOutput", func() {
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

	It("should allow us to produce json output", func() {

		f = func(res pelau.Response) {

			json := &mime.JSONResponse{res}

			json.Send("application/json", map[string]bool{"success": true})

		}

		resp, _ := http.Get(server.URL)
		jmap := map[string]bool{}
		json.NewDecoder(resp.Body).Decode(&jmap)
		Expect(jmap["success"]).To(Equal(true))
	})

})
