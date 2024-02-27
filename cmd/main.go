package main

import (
	"context"
	"fmt"
	"log"
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

	host := config.Host
	port := config.Port
	user := config.User
	password := config.Password
	dbname := config.Dbname

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	postgresAdapter, err := adapters.NewPostgresAdapter(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	r := drivers.SetupRouter(postgresAdapter, firebaseAdapter, postgresAdapter)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
