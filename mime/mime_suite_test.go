package mime_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMid(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mime Suite")
}
