package api

import (
	"fmt"
	"golang-clean-arch/api/routers"
	"golang-clean-arch/config"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		testRouter := v1.Group("/test")

		routers.Health(health)
		routers.TestHandler(testRouter)
	}
	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")

		routers.Health(health)
	}

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		return
	}
}
