package setupteardown

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMeowfacts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Setupteardown Suite")
}

// Run before every container in the test suite
var _ = BeforeSuite(func() {
	fmt.Println("BeforeSuite")
})

// Run after every container in the test suite
var _ = AfterSuite(func() {
	fmt.Println("AfterSuite")
})
