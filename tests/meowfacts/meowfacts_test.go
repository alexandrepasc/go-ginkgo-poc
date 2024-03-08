package meowfacts_test

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/alexandrepasc/go-ginkgo-poc/configurations"
	"github.com/alexandrepasc/go-ginkgo-poc/tests/meowfacts/data"
	"github.com/alexandrepasc/go-ginkgo-poc/tests/meowfacts/types"
	"github.com/gavv/httpexpect"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/xeipuuv/gojsonschema"
)

var _ = Describe("Meowfacts", Label("mewfacts"), func() {
	Describe("Get meowfacts", Label("get"), func() {
		It("No parameters", Label("noparams"), func() {
			req, err := http.NewRequest("GET", "https://meowfacts.herokuapp.com/", http.NoBody)

			// Succeed() simply asserts that a passed-in error is nil
			Expect(err).To(Succeed())

			req.Header.Add("User-Agent", "PostmanRuntime/7.28.4")

			client := &http.Client{}

			resp, err := client.Do(req)

			Expect(err).To(Succeed())

			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			respBody := types.GetMeowfactsResponse{}
			json.NewDecoder(resp.Body).Decode(&respBody)

			Expect(len(respBody.Data)).To(Equal(1))
		})

		Context("With parameter", Label("params"), func() {
			When("Count is", func() {
				It("3", func() {
					req, err := http.NewRequest("GET", "https://meowfacts.herokuapp.com/", http.NoBody)
	
					Expect(err).To(Succeed())
	
					req.Header.Add("User-Agent", "PostmanRuntime/7.28.4")
	
					q := req.URL.Query()
	
					q.Add("count", "3")
	
					req.URL.RawQuery = q.Encode()
	
					client := &http.Client{}
	
					resp, err := client.Do(req)
	
					Expect(err).To(Succeed())
	
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
	
					respBody := types.GetMeowfactsResponse{}
					json.NewDecoder(resp.Body).Decode(&respBody)
	
					Expect(len(respBody.Data)).To(Equal(3))
				})
			})
			
			When("ID is", Label("id"), func() {
				// Test using json schema to validate the response body
				It("1", Label("1"), func() {
					req, err := http.NewRequest("GET", "https://meowfacts.herokuapp.com/", http.NoBody)
	
					Expect(err).To(BeNil())
	
					req.Header.Add("User-Agent", "PostmanRuntime/7.28.4")
	
					q := req.URL.Query()
	
					q.Add("id", "1")
	
					req.URL.RawQuery = q.Encode()
	
					client := &http.Client{}
	
					resp, err := client.Do(req)
	
					Expect(err).To(Succeed())
	
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
	
					e := gojsonschema.NewStringLoader(data.GetMeowfactsId1)
	
					b, err := io.ReadAll(resp.Body)
	
					Expect(err).To(Succeed())
	
					a := gojsonschema.NewBytesLoader(b)
	
					res, err := gojsonschema.Validate(e, a)
	
					Expect(err).To(Succeed())
	
					Expect(res.Valid()).To(BeTrue())
				})
			})

			When("Lang is", Label("lang"), func() {
				// Test using httpexpect package
				It("ger-de", Label("ger", "1"), func() {
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

					resp := a.GET("/").Expect()

					resp.Status(http.StatusOK).
						JSON().Object().
						Schema(data.GetMeowfactsLangGer)
				})
			})
		})
	})
})
