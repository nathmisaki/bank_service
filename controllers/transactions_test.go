package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nelsonmhjr/bank_service/controllers"
	"github.com/nelsonmhjr/bank_service/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transactions", func() {
	AfterEach(func() {
		Db.Exec("Truncate transactions")
	})
	Describe("CreateTransactions POST /transactions", func() {
		var request = func(trxToCreate controllers.TransactionToCreate) *http.Response {
			body, err := json.Marshal(trxToCreate)
			Expect(err).To(BeNil())
			resp, err := http.Post(fmt.Sprintf("%s/transactions/", Ts.URL),
				"application/json", bytes.NewBuffer(body))
			Expect(err).To(BeNil())
			return resp
		}

		Context("without valid input", func() {
			It("should render UnprocessableEntity", func() {
				resp := request(controllers.TransactionToCreate{})
				Expect(resp.StatusCode).To(Equal(422))
			})
		})

		Context("with valid input", func() {
			It("should create Transaction", func() {
				var cntBef, cntAft int
				Db.Model(&models.Transaction{}).Count(&cntBef)
				resp := request(controllers.TransactionToCreate{
					AccountID:       1,
					OperationTypeID: 4,
					Amount:          123.45})

				Expect(resp.StatusCode).To(Equal(201))
				Db.Model(&models.Transaction{}).Count(&cntAft)
				Expect(cntAft).To(Equal(cntBef + 1))
			})
		})
	})
})
