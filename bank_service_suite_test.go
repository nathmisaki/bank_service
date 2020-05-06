package main_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBankService(t *testing.T) {
	RegisterFailHandler(Fail)
	os.Setenv("DB_NAME", "bank_service_test")
	RunSpecs(t, "BankService Suite")
}
