package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/lucky777strike/bottest/handler"
	"github.com/lucky777strike/bottest/repository"
	"github.com/lucky777strike/bottest/usecase"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "127.0.0.1",
		Port:     "5432",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "12345678",
	})
	if err != nil {
		logger.Panicln(err)
	}
	repo := repository.NewRepository(db)
	ucase := usecase.NewService(repo)
	ctx, cancel := context.WithCancel(context.Background())
	token := "5990324330:AAEZdIaNzVTSQIlZJnU9zwj1QhfnPSDXr5g"
	handler := handler.NewHandler(ctx, logger, ucase, token)
	handler.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	logger.Info("Shutting down...")
	cancel()
}
