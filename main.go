package main

import (
	"StatisticColector/dbStats"
	"StatisticColector/endpoints"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

const (
	dbHost = "localhost"
	dbPost = 5432
	dbUser = "user"
	dbName = "statsAndUsers"
	dbPass = "example"
)

func main() {
	log.Println("Connecting to db")

	psqlconn :=
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPost, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err, "fail to connect to db")
	}

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Minute)
	//defer cancel()
	ctx = nil

	Repo := dbStats.NewSQLGORMRepository(ctx, db)
	if err := Repo.Migrate(); err != nil {
		log.Fatalln(err)
	}

	routes := endpoints.NewRouts(Repo)
	routes.AddPath()
	routes.Start(2137)
}
