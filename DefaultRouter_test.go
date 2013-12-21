package pelau_test

import (
	"github.com/metasansana/pelau"
	. "github.com/onsi/ginkgo"
	//	. "github.com/onsi/gomega"
	//	"reflect"
)

var _ = Describe("DefaultRouter", func() {

	var (
		router pelau.Router
		cb     pelau.Callback
	)

	BeforeEach(func() {

		router = pelau.DefaultRouter()

		cb = func(req pelau.Request, res pelau.Response) {

		}

	})

	Describe("the actual routing", func() {

		BeforeEach(func() {

			//		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			//		}))

		})
	})
})
