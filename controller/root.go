package controller

import (
	"StatisticColector/controller/endpoints"
	helthCheck "StatisticColector/controller/helfcheck"
	"github.com/gin-gonic/gin"
	"log"
)

type Routs struct {
	//repo *dbStats.Repo
	r *gin.Engine
}

func NewRouts() *Routs {
	return &Routs{
		//repo: repo,
		r: gin.Default(),
	}
}

func (r *Routs) AddPath() {
	r.r.GET("/", helthCheck.HalfCheck)
	r.r.GET("/tea", helthCheck.Coffee)
	r.r.GET("/coffee", helthCheck.Coffee)

	v1 := r.r.Group("/v1")
	{
		stat := v1.Group("/stat")
		{
			stat.POST("/", endpoints.PostStat)
			stat.GET("/")
			stat.GET("/:id")
			stat.DELETE("/:id")
		}
	}
}

func (r *Routs) Start(port string) {
	portString := ":" + port
	if err := r.r.Run(portString); err != nil {
		log.Fatalln("gin do not start on", port)
	}
}
