package main

import (
	segment "avito-sest-segments"
	"avito-sest-segments/internal/handler"
	"avito-sest-segments/internal/repository"
	"avito-sest-segments/internal/service"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

// @title Avito segment task Api
// @version 1.0
// @description Реализация тестового задания avito go backend-trainee-assignment-2023.

// @host localhost:8080
// @BasePath /

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Error("Config init error: starting")

	if err := InitConfig(); err != nil {
		logrus.Fatalf("Config init error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Usename:  viper.GetString("db.usename"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("DB init fail: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	repository.InitTables(db)
	go func() {
		repository.DeleteExpired(db)
	}()
	srv := new(segment.Server)
	if err := srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Running error: %s", err.Error())
	}

}
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
