package test

import (
	"golang-web-testing/config"
	"golang-web-testing/route"
	"testing"

	"github.com/gin-gonic/gin"
)

var routes *gin.Engine

func TestMain(m *testing.M) {
	routes = route.SetupRoutes()
	config.DBConnect("host=localhost user=postgres password=postgres dbname=webtest-go-example port=5432 sslmode=disable TimeZone=Asia/Jakarta")
	DBCleanup()

	m.Run()
}
