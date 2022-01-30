package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:     viper.GetString("db.dbname"),
		DBPort:     viper.GetString("db.port"),
		SSLMode:    viper.GetString("db.sslmode"),
	}
}

func InitDBConnection(config models.DBConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.SSLMode,
	)
	dbConnection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return dbConnection, err
}

func main() {
	if err := initConfigFile(); err != nil {
		log.Fatalf("Error while init config %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	dbConnection, err := InitDBConnection(initDBConfigFile())
	if err != nil {
		log.Fatalf("Error while connecting to database %s", err.Error())
	}

	storage := storage.NewStorage(dbConnection)
	err = storage.AuthorizationStorage.MigrateTable()
	if err != nil {
		log.Fatalf("Error while migrating database %s", err.Error())
	}
	service := services.NewService(storage)
	authRouter := routers.NewAuthRouter(service)
	authServer := server.NewServer()

	go func() {
		if err := authServer.Run(viper.GetString("authPort"), authRouter.InitAuthRouter()); err != nil {
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
