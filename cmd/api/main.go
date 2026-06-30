package main

import (
	"log"
	"os"

	"Hermes/internal/delivery/http"
	"Hermes/internal/repository"
	"Hermes/internal/repository/data"
	"Hermes/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	log.Print("Start loading")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}
	PORT := os.Getenv("PORT")
	DB_PATH := os.Getenv("DB_PATH")

	dbSetUp := data.SqliteDB{
		DbPath:       DB_PATH,
		MaxOpenConns: 15,
	}

	sqliteDb, err := dbSetUp.Initialize()
	data.Create(sqliteDb)

	if err != nil {
		log.Print("Error unable to connect to the database")
		log.Fatal(err)
	}

	log.Print("User controller set up")
	userRepo := &repository.UserRepository{Db: sqliteDb}
	userUseCase := usecase.NewUserUseCase(userRepo)

	g := gin.Default()

	http.RegisterUserRoutes(g, userUseCase)

	g.Run(":" + PORT)
}
