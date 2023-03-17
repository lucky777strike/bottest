package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lucky777strike/bottest/handler"
	"github.com/lucky777strike/bottest/repository"
	"github.com/lucky777strike/bottest/usecase"
	"github.com/sirupsen/logrus"
)

func main() {

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "127.0.0.1",
		Port:     "5432",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "12345678",
	})
	if err != nil {
		log.Panicln(err)
	}
	StatRep := repository.NewStatisticsPostgresRepository(db)
	WeatherRepo := repository.NewWeatherPostgresRepository(db)
	ucase := usecase.NewService(StatRep, WeatherRepo)
	ctx, cancel := context.WithCancel(context.Background())
	token := "5990324330:AAEZdIaNzVTSQIlZJnU9zwj1QhfnPSDXr5g"
	logger := logrus.New()
	handler := handler.NewHandler(ctx, cancel, logger, ucase, token)
	handler.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	logger.Info("Shutting down...")
	cancel()
}

func createLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logger.SetLevel(logrus.InfoLevel)

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(io.MultiWriter(os.Stdout, file))
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}

	return logger
}
