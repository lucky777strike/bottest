package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/lucky777strike/bottest/handler"
	"github.com/lucky777strike/bottest/repository"
	"github.com/lucky777strike/bottest/usecase"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logger := logrus.New()
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	//use env for sensetive data
	if err := godotenv.Load(); err != nil {
		logger.Info("no env file: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logger.Fatal("err to init db")
	}
	repo := repository.NewRepository(db)
	ucase := usecase.NewService(repo)
	ctx, cancel := context.WithCancel(context.Background())
	token := os.Getenv("TG_TOKEN")
	handler := handler.NewHandler(ctx, logger, ucase, token)
	handler.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	logger.Info("Shutting down...")
	cancel()
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
