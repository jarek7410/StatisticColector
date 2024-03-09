package endpoints

import (
	"StatisticColector/dbStats"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type Routs struct {
	repo *dbStats.Repo
	r    *gin.Engine
}

func NewRouts(repo *dbStats.Repo) *Routs {
	return &Routs{
		repo: repo,
		r:    gin.Default(),
	}
}

func (r *Routs) AddPath() {
	base := NewBasicForRouts(r.repo)

	r.r.GET("/", halfCheck)
	r.r.GET("/tea", coffee)
	r.r.GET("/coffee", coffee)
	v1 := r.r.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/", base.postUser)
			user.GET("/:id", base.getUserById)
			user.GET("/", base.getUsers)
		}
		stat := v1.Group("/stat")
		{
			stat.POST("/", base.postStat)
			stat.GET("/user/:id", base.getStats)
			stat.GET("/:id", getStatById)
			stat.DELETE("/:id", deleteStat)
		}
	}
}

func (r *Routs) Start(port int16) {
	portString := ":" + strconv.Itoa(int(port))
	if err := r.r.Run(portString); err != nil {
		log.Fatalln("gin do not start on", port)
	}
}
