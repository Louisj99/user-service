package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"user-service/pkg/adapters"
	"user-service/pkg/drivers"
)

func main() {

	config, err := adapters.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	firebaseAdapter, err := adapters.NewFirebaseAdapter(ctx, config.PathToServiceAccountKey)
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	postgresAdapter, err := adapters.NewPostgresAdapter(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	r := drivers.SetupRouter(firebaseAdapter, postgresAdapter)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
