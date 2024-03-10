package catfacts

import (
	"net/http"
	"testing"

	"github.com/alexandrepasc/go-ginkgo-poc/configurations"
	"github.com/alexandrepasc/go-ginkgo-poc/tests/catfacts/data"
	"github.com/gavv/httpexpect"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMeowfacts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Catfacts Suite")
}

var _ = Describe("Catfacts", Label("catfacts"), func() {
	It("Invalid methods", Label("invalid"), func() {
		e := httpexpect.WithConfig(httpexpect.Config{
			BaseURL:  configurations.Endpoints.Catfacts.BaseURL,
			Reporter: httpexpect.NewRequireReporter(GinkgoT()),
			Printers: []httpexpect.Printer{
				httpexpect.NewCurlPrinter(GinkgoT()),
				httpexpect.NewDebugPrinter(GinkgoT(), true),
			},
		})

		a := e.Builder(func(r *httpexpect.Request) {
			r.WithHeaders(map[string]string{"User-Agent": "PostmanRuntime/7.28.4"})
		})

		resp := a.POST("/").Expect()

		resp.Status(http.StatusUnauthorized)

		resp = a.PUT("/").Expect()

		resp.Status(http.StatusNotFound)

		resp = a.PATCH("/").Expect()

		resp.Status(http.StatusNotFound)

		resp = a.DELETE("/").Expect()

		resp.Status(http.StatusNotFound)
	})

	Describe("Get", Label("get"), func() {
		It("List", func() {
			e := httpexpect.WithConfig(httpexpect.Config{
				BaseURL:  configurations.Endpoints.Catfacts.BaseURL,
				Reporter: httpexpect.NewRequireReporter(GinkgoT()),
				Printers: []httpexpect.Printer{
					httpexpect.NewCurlPrinter(GinkgoT()),
					httpexpect.NewDebugPrinter(GinkgoT(), true),
				},
			})
	
			a := e.Builder(func(r *httpexpect.Request) {
				r.WithHeaders(map[string]string{"User-Agent": "PostmanRuntime/7.28.4"})
			})

			resp := a.GET("/").Expect()

			resp.Status(http.StatusOK).
				JSON().Array().
				Schema(data.GetCatfactsList)
		})

		It("Id", func() {
			e := httpexpect.WithConfig(httpexpect.Config{
				BaseURL:  configurations.Endpoints.Catfacts.BaseURL,
				Reporter: httpexpect.NewRequireReporter(GinkgoT()),
				Printers: []httpexpect.Printer{
					httpexpect.NewCurlPrinter(GinkgoT()),
					httpexpect.NewDebugPrinter(GinkgoT(), true),
				},
			})
	
			a := e.Builder(func(r *httpexpect.Request) {
				r.WithHeaders(map[string]string{"User-Agent": "PostmanRuntime/7.28.4"})
			})

			const id = "58e008780aac31001185ed05"

			resp := a.GET("/" + id).Expect()

			resp.Status(http.StatusOK).
				JSON().Object().
				Schema(data.GetCatfactsId)
		})
	})
})
