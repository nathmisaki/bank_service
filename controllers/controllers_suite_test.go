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

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

func SetupServer() (*httptest.Server, *gorm.DB) {
	gin.SetMode(gin.TestMode)
	db := models.SetupModels(gin.Mode())
	ts := httptest.NewServer(routes.SetupRouter(db))
	return ts, db
}
