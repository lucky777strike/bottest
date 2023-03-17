package main

import (
	"context"
	"log"

	"github.com/lucky777strike/bottest/handler"
	"github.com/lucky777strike/bottest/repository"
	"github.com/lucky777strike/bottest/usecase"
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
	ucase := usecase.NewService(StatRep)
	ctx, cancel := context.WithCancel(context.Background())
	token := "5990324330:AAEZdIaNzVTSQIlZJnU9zwj1QhfnPSDXr5g"
	handler := handler.NewHandler(ctx, cancel, ucase, token)
	handler.Start()

	for {
	}
}
