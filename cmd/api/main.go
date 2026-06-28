package main

import (
	"log"
	"os"

	"Hermes/internal/delivery/http"
	"Hermes/internal/repository/data"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

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
	if err != nil {
		log.Fatal("Error unable to connect to the database")
		log.Fatal(err)
	}

	g := gin.Default()

	http.RegisterUserRoutes(g, sqliteDb)

	g.Run(PORT)
}
