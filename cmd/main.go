package main

import (
	"context"
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

	r := drivers.SetupRouter(firebaseAdapter)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
