package main

import (
	"StatisticColector/controller"
	"StatisticColector/database"
	"StatisticColector/model"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	log.Println("Connecting to db")
	loadEnv()

	initdatabase()

	serveApplication()
}

func initdatabase() {
	database.InitDatabase()
	err := database.Re.DB.AutoMigrate(
		&model.User{},
		&model.Name{},
		&model.Stat{},
	)
	if err != nil {
		log.Println(`automation error:`, err)
	}
}

func serveApplication() {

	routes := controller.NewRouts()
	routes.AddPath()
	routes.Start(os.Getenv("PORT"))
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {

		log.Println("Error loading .env file")
		shell := os.Getenv("SHELL")
		log.Println(shell)

	}
	log.Println(".env file loaded successfully")
}
