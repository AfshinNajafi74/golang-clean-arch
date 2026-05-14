package api

import (
	"fmt"
	"golang-clean-arch/api/routers"
	"golang-clean-arch/api/validations"
	"golang-clean-arch/config"
	"golang-clean-arch/docs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		if err != nil {
			return
		}
	}
	r.Use(
		gin.Logger(),
		//gin.Recovery(),
		//middlewares.TestMiddlewares()
	)

	RegisterSwagger(r, cfg)

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

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = " golang web api"
	docs.SwaggerInfo.Description = "This is a web api server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
