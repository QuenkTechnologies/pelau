package mime_test

import (
	"bytes"
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
		f      func(pelau.Request)
		result map[string]bool
		data   map[string]bool
	)

	BeforeEach(func() {

		data = map[string]bool{"success": true}

		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			f(pelau.DefaultRequest(&pelau.ModifiedRequest{[]string{}, r}))

		}))

	})

	AfterEach(func() {

		defer server.Close()

	})

	It("should allow us to produce json output", func() {

		f = func(req pelau.Request) {

			json := &mime.JSONRequest{req}
			json.Decode("application/json", &result)
			Expect(result["success"]).To(Equal(true))

		}

		client := &http.Client{}
		toSend, _ := json.Marshal(data)
		creq, _ := http.NewRequest("POST", server.URL, bytes.NewReader(toSend))
		creq.Header.Add("Content-Type", "application/json")
		resp, _ := client.Do(creq)
		defer resp.Body.Close()
	})

})
