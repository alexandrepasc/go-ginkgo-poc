package meowfacts_test

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/alexandrepasc/go-ginkgo-poc/tests/meowfacts/data"
	"github.com/alexandrepasc/go-ginkgo-poc/tests/meowfacts/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xeipuuv/gojsonschema"
)

var _ = Describe("Meowfacts", func() {
	Describe("Get meowfacts", func() {
		It("No parameters", func() {
			req, err := http.NewRequest("GET", "https://meowfacts.herokuapp.com/", http.NoBody)

			Expect(err).To(BeNil())

			req.Header.Add("User-Agent", "PostmanRuntime/7.28.4")

			client := &http.Client{}

			resp, err := client.Do(req)

			Expect(err).To(BeNil())

			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			respBody := types.GetMeowfactsResponse{}
			json.NewDecoder(resp.Body).Decode(&respBody)

			Expect(len(respBody.Data)).To(Equal(1))
		})

		Context("With parameter", func() {
			When("Count is", func() {
				It("3", func() {
					req, err := http.NewRequest("GET", "https://meowfacts.herokuapp.com/", http.NoBody)
	
					Expect(err).To(BeNil())
	
					req.Header.Add("User-Agent", "PostmanRuntime/7.28.4")
	
					q := req.URL.Query()
	
					q.Add("count", "3")
	
					req.URL.RawQuery = q.Encode()
	
					client := &http.Client{}
	
					resp, err := client.Do(req)
	
					Expect(err).To(BeNil())
	
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
	
					respBody := types.GetMeowfactsResponse{}
					json.NewDecoder(resp.Body).Decode(&respBody)
	
					Expect(len(respBody.Data)).To(Equal(3))
				})
			})
			
			When("ID is", func() {
				// Test using json schema to validate the response body
				It("1", func() {
					req, err := http.NewRequest("GET", "https://meowfacts.herokuapp.com/", http.NoBody)
	
					Expect(err).To(BeNil())
	
					req.Header.Add("User-Agent", "PostmanRuntime/7.28.4")
	
					q := req.URL.Query()
	
					q.Add("id", "1")
	
					req.URL.RawQuery = q.Encode()
	
					client := &http.Client{}
	
					resp, err := client.Do(req)
	
					Expect(err).To(BeNil())
	
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
	
					e := gojsonschema.NewStringLoader(data.GetMeowfactsId1)
	
					b, err := io.ReadAll(resp.Body)
	
					Expect(err).To(BeNil())
	
					a := gojsonschema.NewBytesLoader(b)
	
					res, err := gojsonschema.Validate(e, a)
	
					Expect(err).To(BeNil())
	
					Expect(res.Valid()).To(BeTrue())
				})
			})
		})
	})
})
