package endpoints

import (
	"StatisticColector/controller/endpoints/dto"
	"StatisticColector/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostStat(c *gin.Context) {
	postStat := &dto.PostStatDto{}

	if err := c.ShouldBindJSON(&postStat); err != nil {
		c.JSON(http.StatusBadRequest, postStat)
		return
	}
	if postStat.Name == "" && postStat.NameID == 0 {
		c.Status(http.StatusBadRequest)
		return
	}
	if postStat.NameID == 0 {
		name := &model.Name{Name: postStat.Name}
		name.GetByName()
		if name.ID == 0 {
			//log.Println("no record fide with name:", name.Name, name.ID)
			name.Save()
		}
		//log.Println("record ID:", name.ID)
		postStat.NameID = name.ID
	}
	stat := &model.Stat{
		Value:  postStat.Value,
		Type:   postStat.Type,
		NameID: postStat.NameID,
	}
	stat.Save()
	c.JSONP(http.StatusOK, stat)
}

func GetStats(context *gin.Context) {

}
