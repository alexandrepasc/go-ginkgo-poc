package meowfacts_test

import (
	"net/http"

	"github.com/alexandrepasc/go-ginkgo-poc/configurations"
	"github.com/gavv/httpexpect"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Meowfacts Negative", Label("meowfacts negative"), func() {
	It("Post meowfacts", Label("post"), func () {
		e := httpexpect.WithConfig(httpexpect.Config{
			BaseURL:  configurations.Endpoints.BaseURL,
			Reporter: httpexpect.NewRequireReporter(GinkgoT()),
			Printers: []httpexpect.Printer{
				httpexpect.NewCurlPrinter(GinkgoT()),
				httpexpect.NewDebugPrinter(GinkgoT(), true),
			},
		})

		a := e.Builder(func(r *httpexpect.Request) {
			r.WithQueryObject(map[string]interface{}{"lang": "ger-de", "id": "1"})
			r.WithHeaders(map[string]string{"User-Agent": "PostmanRuntime/7.28.4"})
		})

		resp := a.POST("/").Expect()

		resp.Status(http.StatusBadRequest)
	})
})
