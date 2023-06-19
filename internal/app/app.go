package app

import (
	"AuthApp/internal/config"
	"AuthApp/internal/delivery/handler"
	"AuthApp/internal/repository"
	"AuthApp/internal/server"
	"AuthApp/internal/service"
	"AuthApp/pkg/database/postgres"
	"AuthApp/pkg/logger"
)

func Run(configPath string) {
	configs, error := config.Init(configPath)
	if error != nil {
		logger.Error(error.Error())
		return
	}

	db, error := postgres.Init(configs.Postgres)
	if error != nil {
		logger.Error(error.Error())
		return
	}

	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories)
	handlers := handler.Init(services)

	srv := server.Init(configs.Http, handlers.InitRoutes())
	if error := srv.Run(); error != nil {
		logger.Error(error.Error())
		return
	}
}
