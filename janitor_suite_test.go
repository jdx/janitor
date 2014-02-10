package janitor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestJanitor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Janitor Suite")
}
