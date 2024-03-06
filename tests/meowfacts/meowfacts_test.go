package meowfacts_test

import (
	"encoding/json"
	"net/http"

	"github.com/alexandrepasc/go-ginkgo-poc/tests/meowfacts/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

			respBody := models.GetMeowfactsResponse{}
			json.NewDecoder(resp.Body).Decode(&respBody)

			Expect(len(respBody.Data)).To(Equal(1))
		})

		Describe("With parameters", func() {
			It("Count 3", func() {
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

				respBody := models.GetMeowfactsResponse{}
				json.NewDecoder(resp.Body).Decode(&respBody)

				Expect(len(respBody.Data)).To(Equal(3))
			})
		})
	})
})
