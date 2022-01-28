package main

import (
	"context"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	server "pmokeev/web-chat/internal"
	"pmokeev/web-chat/internal/models"
	"pmokeev/web-chat/internal/routers"
	"pmokeev/web-chat/internal/services"
	"pmokeev/web-chat/internal/storage"
	"syscall"
)

func initConfigFile() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDBConfigFile() models.DBConfig {
	return models.DBConfig{
		DBHost:     viper.GetString("db.host"),
		DBUser:     viper.GetString("db.user"),
		DBPassword: viper.GetString("db.password"),
		DBName:     viper.GetString("db.dbname"),
		DBPort:     viper.GetString("db.port"),
		SSLMode:    viper.GetString("db.sslmode"),
	}
}

func main() {
	if err := initConfigFile(); err != nil {
		log.Fatalf("Error while init config %s", err.Error())
	}

	authStorage := storage.NewAuthStorage()
	if err := authStorage.InitDBConnection(initDBConfigFile()); err != nil {
		log.Fatalf("Error while connecting to database %s", err.Error())
	}
	authService := services.NewAuthService(authStorage)
	router := routers.NewAuthRouter(authService)
	authServer := server.NewServer()

	go func() {
		if err := authServer.Run(viper.GetString("authPort"), router.InitAuthRouter()); err != nil {
			log.Fatalf("Error while running server %s", err.Error())
		}
	}()

	log.Print("API started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	log.Print("API shutdowned")

	if err := authServer.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error while shutdowning server %s", err.Error())
	}
}
