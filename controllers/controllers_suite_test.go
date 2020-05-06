package controllers_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nelsonmhjr/bank_service/models"
	"github.com/nelsonmhjr/bank_service/routes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var Db *gorm.DB
var Ts *httptest.Server

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = BeforeSuite(func() {
	gin.SetMode(gin.TestMode)
	Db = models.SetupModels(gin.Mode())
	Ts = httptest.NewServer(routes.SetupRouter(Db))
})
