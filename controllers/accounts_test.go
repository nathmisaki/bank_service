package controllers_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

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
	AfterEach(func() {
		Db.Exec("Truncate bank_accounts")
	})

	Describe("FindAccount GET /accounts/:accountId", func() {
		var request = func(id uint) *http.Response {
			resp, err := http.Get(fmt.Sprintf("%s/accounts/%d", Ts.URL, id))
			Expect(err).To(BeNil())
			return resp
		}

		var requestBody = func(resp *http.Response) string {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			Expect(err).To(BeNil())
			return string(bodyBytes)
		}

		Context("without a matching BankAccount with the matchin ID", func() {
			It("should render a Bad Request", func() {
				resp := request(1)
				Expect(resp.StatusCode).To(Equal(400))
			})
		})

		Context("with a matching BankAccount ID", func() {
			It("should render JSON with id and document_number", func() {
				bankAcc := models.BankAccount{DocumentNumber: "123456"}
				Db.Create(&bankAcc)
				resp := request(bankAcc.ID)
				bodyString := requestBody(resp)
				Expect(bodyString).To(
					MatchJSON(
						fmt.Sprintf(`{"document_number":"%s","account_id":%d}`,
							bankAcc.DocumentNumber, bankAcc.ID)))
			})
		})
	})

	Describe("CreateAccounts POST /accounts", func() {
		Context("without a json body", func() {
			It("should render a UnprocessableEntity Status Code", func() {
				// resp, err := http.Get(fmt.Sprintf("%s/accounts/", Ts.URL))

			})
		})
	})
})
