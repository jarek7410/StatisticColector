package controller

import (
	"StatisticColector/dto"
	model "StatisticColector/model/helthcheck"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HalfCheck godoc
// @Title halfCheck
// @Summary Get status about server helth
// @Description  get status
// @Produce json
// @Success 200 {object} dto.HalfCheckDto
// @Router /half_check [GET]
func HalfCheck(c *gin.Context) {
	va := model.HealthCheck()
	var dbStatus string
	if va != nil {
		dbStatus = "down"
	} else {
		dbStatus = "up"
	}
	//data := map[string]interface{}{
	data := dto.HalfCheckDto{
		Version:     "0.1",
		VersionName: "",
		Database:    dbStatus,
	}
	c.JSONP(http.StatusOK, data)
}
func Coffee(c *gin.Context) {
	//c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	c.String(http.StatusTeapot, "I am a teapot but programmer need coffee!!!\nblik +48 735 917 147")
}
