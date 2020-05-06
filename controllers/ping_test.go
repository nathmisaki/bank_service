package controllers_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ping", func() {
	Context("GET /ping", func() {
		It("Should respond 'pong'", func() {
			// Make a request to our server with the {base url}/ping
			resp, err := http.Get(fmt.Sprintf("%s/ping", Ts.URL))
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			Expect(err).To(BeNil())
			bodyString := string(bodyBytes)
			Expect(bodyString).To(Equal("pong"))
		})
	})

})
