package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"register-profile-service/internal/handler"
	"register-profile-service/internal/models"
	"register-profile-service/internal/repository"
	"register-profile-service/internal/service"
	"register-profile-service/internal/service/hash"
	"register-profile-service/internal/service/token"
	"register-profile-service/pkg/logging"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("../configs/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
}
func main() {
	if err := godotenv.Load(); err != nil {
		logging.Logger.Warn(logging.MakeLog("failed to initialization DB ", err))
		return
	}
	if err := logging.NewLogService(os.Stdout, os.Getenv("LOG_MODE")); err != nil {
		log.Fatal("Failed to initialize logger: ", err)
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
		logging.Logger.Warn(logging.MakeLog("failed to initialization DB ", err))
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		logging.Logger.Warn(logging.MakeLog("JWT_SECRET must be set", err))
		return
	}
	tokenMaker, err := token.NewJWTMaker(secret)
	if err != nil {
		logging.Logger.Warn(logging.MakeLog("failed to initialization tokenMarker", err))
		return
	}
	hashUtil := hash.NewBcryptHasher()

	ctx := context.Background()
	repository := repository.NewRepository(ctx, db)
	logging.Logger.Info("initializing repository")
	services := service.NewService(ctx, repository, tokenMaker, hashUtil)
	logging.Logger.Info("initializing services")
	handlers := handler.NewHandler(services)
	logging.Logger.Info("initializing routes")
	serv := new(models.Server)
	go func() {
		if err := serv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
			logging.Logger.Warn(logging.MakeLog("error starting server", err))
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serv.Close(ctx); err != nil {
		logging.Logger.Warn(logging.MakeLog("error starting server", err))
	}

	logging.Logger.Info("Server start on port: " + viper.GetString("port"))

}
