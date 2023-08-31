package main

import (
	"avito-sest-segments/internal/repository"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

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
	//TODO: Init server

}
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
