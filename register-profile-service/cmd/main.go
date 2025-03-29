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

}
