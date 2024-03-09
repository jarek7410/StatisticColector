package endpoints

import (
	"StatisticColector/dbStats"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Base struct {
	repo *dbStats.Repo
}

func NewBasicForRouts(repo *dbStats.Repo) *Base {
	return &Base{
		repo: repo,
	}
}

func halfCheck(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}
func coffee(c *gin.Context) {
	c.String(http.StatusTeapot, "I am a teapot but programer need coffee!!!")
}
func (r *Base) postUser(c *gin.Context) {
	var user dbStats.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		log.Println("some error", err)
		return
	}
	createUser, err := r.repo.CrateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data, database error"})
		log.Println("some error", err)
		return
	}
	c.JSON(http.StatusOK, createUser)
}
func (r *Base) getUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not integer"})
		return
	}

	user, err := r.repo.GerUser(uint(userId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)

}
func (r *Base) getUsers(c *gin.Context) {
	users, err := r.repo.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no users or not found"})
		return
	}
	c.JSON(http.StatusOK, users)
}
func (r *Base) postStat(c *gin.Context) {
	var stat dbStats.Stat
	if err := c.ShouldBindJSON(&stat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		log.Println("some error", err.Error())
	}

	//userId, err := strconv.Atoi(c.Param("userId"))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "id is not integer"})
	//	return
	//}
	//user, err1 := r.repo.GerUser(uint(userId))
	user, err1 := r.repo.GerUser(stat.UserID)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong id"})
		return
	}
	createStat, err2 := r.repo.CrateStat(*user, stat)
	if err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	}
	c.JSON(http.StatusOK, createStat)

}
func (r *Base) getStats(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not integer"})
		return
	}

	user, err := r.repo.GerUser(uint(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	stats, err := r.repo.GetStats(*user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "stats not found"})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func getStatById(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
func deleteStat(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
