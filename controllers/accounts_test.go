package controllers_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nelsonmhjr/bank_service/models"
	"github.com/nelsonmhjr/bank_service/routes"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var _ = Describe("Accounts", func() {

	Describe("FindAccount GET /accounts/:accountId", func() {
		Context("Without a matching BankAccount with the matchin ID", func() {
			It("Should render a Bad Request", func() {
				db := models.SetupModels()
				r := routes.SetupRouter(db)

				w := performRequest(r, "GET", "/accounts/1")

				Expect(w.Code).To(Equal(400))
			})
		})
	})
})
