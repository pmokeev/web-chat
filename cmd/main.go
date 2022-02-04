package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	server "pmokeev/web-chat/internal"
	"pmokeev/web-chat/internal/models"
	"pmokeev/web-chat/internal/routers"
	"pmokeev/web-chat/internal/services"
	"pmokeev/web-chat/internal/storage"
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

func InitAuthServerConfiguration(storage *storage.Storage) *routers.AuthRouter {
	err := storage.AuthorizationStorage.MigrateTable()
	if err != nil {
		log.Fatalf("Error while migrating database %s", err.Error())
	}
	service := services.NewService(storage)
	authRouter := routers.NewAuthRouter(service)

	return authRouter
}

func InitChatServerConfiguration(storage *storage.Storage) *routers.ChatRouter {
	err := storage.AuthorizationStorage.MigrateTable()
	if err != nil {
		log.Fatalf("Error while migrating database %s", err.Error())
	}
	service := services.NewService(storage)
	chatRouter := routers.NewChatRouter(service)

	return chatRouter
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

	mainStorage := storage.NewStorage(dbConnection)
	authRouter := InitAuthServerConfiguration(mainStorage)
	chatRouter := InitChatServerConfiguration(mainStorage)
	authServer := server.NewServer()
	charServer := server.NewServer()

	var serverGroup errgroup.Group
	serverGroup.Go(func() error {
		return authServer.Run(viper.GetString("authPort"), authRouter.InitAuthRouter())
	})
	serverGroup.Go(func() error {
		return charServer.Run(viper.GetString("chatPort"), chatRouter.InitChatRouter())
	})

	if err := serverGroup.Wait(); err != nil {
		log.Fatal(err)
	}
}
