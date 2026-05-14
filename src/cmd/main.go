package main

import (
	"golang-clean-arch/api"
	"golang-clean-arch/config"
	"golang-clean-arch/data/cache"
)

func main() {
	cfg := config.GetConfig()
	api.InitServer(cfg)
	defer cache.CloseRedis()
	cache.InitRedis(cfg)
}
