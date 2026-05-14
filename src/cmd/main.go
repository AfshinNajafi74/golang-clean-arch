package main

import (
	"golang-clean-arch/api"
	"golang-clean-arch/config"
	"golang-clean-arch/data/cache"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	api.InitServer(cfg)
	//logger := logging.NewLogger(cfg)
	defer cache.CloseRedis()
	cache.InitRedis(cfg)
}
