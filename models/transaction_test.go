package models_test

import (
	"github.com/nelsonmhjr/bank_service/controllers"
	"github.com/nelsonmhjr/bank_service/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transaction", func() {
	Describe("ValidateAndCreate", func() {
		Context("with empty struct", func() {
			It("should panic", func() {
				Expect(func() {
					models.Transaction{}.
						ValidateAndCreate(&controllers.TransactionToCreate{})
				}).To(Panic())
			})
		})
	})
})
