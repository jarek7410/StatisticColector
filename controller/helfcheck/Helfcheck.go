package controller

import (
	model "StatisticColector/model/helthcheck"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin-swagger middleware
// swagger embed files

// @BasePath /
// helth check for basik working of this app
// @Summary blablabla
func HalfCheck(c *gin.Context) {
	va := model.HealthCheck()
	var dbStatus string
	if va != nil {
		dbStatus = "down"
	} else {
		dbStatus = "up"
	}
	data := map[string]interface{}{
		"version":  "alpha",
		"database": dbStatus,
	}
	c.JSONP(http.StatusOK, data)
}
func Coffee(c *gin.Context) {
	c.String(http.StatusTeapot, "I am a teapot but programer need coffee!!!\nblik 735917147")
}
