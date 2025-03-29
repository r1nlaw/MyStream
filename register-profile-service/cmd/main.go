package main

import (
	"context"
	"fmt"
	"os"
	"register-profile-service/internal/repository"
	"register-profile-service/pkg/logging"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Errorf(logging.MakeLog("Failed read .env file", err))
	}
	logging.NewLogService(os.Stdout, os.Getenv("LOG_MODE"))
	logging.Logger.Debug(logging.MakeLog("Loading configs", nil))

	db, err := repository.InitDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logging.Logger.Warn(logging.MakeLog("Failed to initialization DB ", err))
	}
	ctx := context.Background()
	repository := repository.NewRepository(ctx, db)
	logging.Logger.Info("initializing repository")
	services := service.NewService(repository, ctx)
	logging.Logger.Info("Инициалиазация сервисов")

}
