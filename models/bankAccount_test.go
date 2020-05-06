package models_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	"github.com/nelsonmhjr/bank_service/mocks"
	"github.com/nelsonmhjr/bank_service/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BankAccount", func() {
	Describe("ValidateAndCreate", func() {
		mockCtrl := gomock.NewController(GinkgoT())
		defer mockCtrl.Finish()
		dummyError := errors.New("dummy error")
		var mockBindable *mocks.MockBindable

		Context("with invalid bindable data", func() {
			It("should Return Error", func() {
				mockBindable = mocks.NewMockBindable(mockCtrl)
				accToCreate := models.AccountToCreate{}
				mockBindable.EXPECT().ShouldBind(&accToCreate).Return(dummyError)
				bacc := models.BankAccount{}
				err := bacc.ValidateAndCreate(Db, mockBindable)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
