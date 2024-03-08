package meowfacts_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMeowfacts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Meowfacts Suite")
}
