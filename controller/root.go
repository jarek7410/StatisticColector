package controller

import (
	"StatisticColector/controller/endpoints"
	helthCheck "StatisticColector/controller/helfcheck"
	docs "StatisticColector/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.r.GET("/half_check", helthCheck.HalfCheck)
	r.r.GET("/tea", helthCheck.Coffee)
	r.r.GET("/coffee", helthCheck.Coffee)
	docs.SwaggerInfo.BasePath = "/"
	r.r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.r.Group("/v1")
	{
		stat := v1.Group("/stat")
		{
			stat.POST("/", endpoints.PostStat)
			stat.DELETE("/:id")
		}
		name := v1.Group("/name")
		{
			stat.GET("/:id", endpoints.GetStats)
			name.GET("/", endpoints.GetNames)
		}
	}
}

func (r *Routs) Start(port string) {
	portString := ":" + port
	if err := r.r.Run(portString); err != nil {
		log.Fatalln("gin do not start on", port)
	}
}
