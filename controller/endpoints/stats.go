package endpoints

import (
	"StatisticColector/controller/endpoints/dto"
	"StatisticColector/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func GetStats(c *gin.Context) {
	id, err1 := strconv.ParseUint(c.Param("id"), 10, 0)
	limit, err2 := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 0)
	offset, err3 := strconv.ParseUint(c.DefaultQuery("offset", "0"), 10, 0)
	if err1 != nil || err2 != nil || err3 != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	name := &model.Name{}
	name.ID = uint(id)
	name.GetStats(uint(limit), uint(offset))
	var stats []dto.StatDto
	for _, stat := range name.Stats {
		s := dto.StatDto{
			Value:     stat.Value,
			Type:      stat.Type,
			CreatedAt: stat.CreatedAt,
		}
		stats = append(stats, s)
	}
	res := &dto.GetStatsDto{
		Limit:  int(limit),
		Offset: int(offset),
		Name:   name.Name,
		NameID: name.ID,
		Stats:  stats,
	}
	c.JSON(http.StatusOK, res)
}

func GetNames(c *gin.Context) {
	limit, err2 := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 0)
	offset, err3 := strconv.ParseUint(c.DefaultQuery("offset", "0"), 10, 0)
	if err2 != nil || err3 != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, model.GetAllNames(uint(limit), uint(offset)))
}
