package controllers_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/jinzhu/gorm"
	"github.com/nelsonmhjr/bank_service/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var _ = Describe("Accounts", func() {
	var ts *httptest.Server
	var db *gorm.DB
	BeforeEach(func() {
		ts, db = SetupServer()
	})

	AfterEach(func() {
		db.Exec("Truncate bank_accounts")
	})

	Describe("FindAccount GET /accounts/:accountId", func() {
		Context("Without a matching BankAccount with the matchin ID", func() {
			It("Should render a Bad Request", func() {
				resp, err := http.Get(fmt.Sprintf("%s/accounts/1", ts.URL))
				Expect(err).To(BeNil())

				Expect(resp.StatusCode).To(Equal(400))
			})
		})

		Context("With a matching BankAccount ID", func() {
			It("Should render JSON with id and document_number", func() {
				bankAcc := models.BankAccount{DocumentNumber: "123456"}
				db.Create(&bankAcc)
				resp, err := http.Get(fmt.Sprintf("%s/accounts/%d", ts.URL, bankAcc.ID))
				Expect(err).To(BeNil())
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				Expect(err).To(BeNil())
				bodyString := string(bodyBytes)
				Expect(bodyString).To(MatchJSON(fmt.Sprintf(`{"document_number":"%s","account_id":%d}`, bankAcc.DocumentNumber, bankAcc.ID)))
			})
		})
	})
})
