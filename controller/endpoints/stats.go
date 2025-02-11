package endpoints

import (
	"StatisticColector/dto"
	"StatisticColector/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// PostStat godoc
// @Title post stat
// @Summary create new stat
// @Description  crate new stat
// @Produce json
// @Accept json
// @Param stat body dto.PostStatDto true "new stat"
// @Success 200 {object} model.Stat
// @Failure 400 {object} dto.PostStatDto
// @Failure 400
// @Router /v1/stat [POST]
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

// GetStats godoc
// @Title get stat
// @Summary get existing stats
// @Description  get existing stats by name id
// @Produce json
// @Param id    path int true "id of name of stats which you want"
// @Param limit query int false " "  minimum(1) maximum(100)
// @Param offset query int false " " minimum(0)
// @Success 200 {object} dto.GetStatsDto
// @Failure 400
// @Router /v1/name/{id} [GET]
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

// GetNames godoc
// @Title get names
// @Summary get existing names
// @Description  get existing names in service
// @Produce json
// @Param limit query int false " "  minimum(1) maximum(100)
// @Param offset query int false " " minimum(0)
// @Success 200 {array} dto.NameDto
// @Failure 400
// @Router /v1/name/ [GET]
func GetNames(c *gin.Context) {
	limit, err2 := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 0)
	offset, err3 := strconv.ParseUint(c.DefaultQuery("offset", "0"), 10, 0)
	if err2 != nil || err3 != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, model.GetAllNames(uint(limit), uint(offset)))
}
