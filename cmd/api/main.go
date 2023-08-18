package main

import (
	"github.com/quangdangfit/gocommon/logger"
	"github.com/quangdangfit/gocommon/redis"
	"github.com/quangdangfit/gocommon/validation"

	"goshop/app/dbs"
	"goshop/config"
	httpServer "goshop/internal/server/http"
)

//	@title			GoShop Swagger API
//	@version		1.0
//	@description	Swagger API for GoShop.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Quang Dang
//	@contact.email	quangdangfit@gmail.com

//	@license.name	MIT
//	@license.url	https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//	@BasePath	/api/v1

func main() {
	cfg := config.GetConfig()
	logger.Initialize(cfg.Environment)
	dbs.Init()

	validator := validation.New()

	cache := redis.New(redis.Config{
		Address:  cfg.RedisURI,
		Password: cfg.RedisPassword,
		Database: cfg.RedisDB,
	})

	server := httpServer.NewServer()
	if err := server.Run(validator, cache); err != nil {
		logger.Fatal(err)
	}
}
