package models_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nelsonmhjr/bank_service/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var Db *gorm.DB

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}

var _ = BeforeSuite(func() {
	gin.SetMode(gin.TestMode)
	Db = models.SetupModels(gin.Mode())
})
