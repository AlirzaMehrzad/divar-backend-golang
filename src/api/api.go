package api

import (
	"fmt"

	"github.com/AlirzaMehrzad/divar-golang/src/api/routers"
	"github.com/AlirzaMehrzad/divar-golang/src/configs"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := configs.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
