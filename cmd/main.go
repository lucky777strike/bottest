package main

import (
	"log"
	"os"

	"github.com/lucky777strike/bottest/repository"
	"github.com/lucky777strike/bottest/usecase"
)

func main() {

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "db",
		Port:     "5432",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Panicln()
	}
	rep := repository.NewRepository(db)
	ucase := usecase.NewUsecase(rep)
}
