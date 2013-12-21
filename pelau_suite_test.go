package pelau_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPelau(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pelau Suite")
}
