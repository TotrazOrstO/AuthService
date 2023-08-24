package main

import (
	"MedodsProject/internal/delivery/http"
	"MedodsProject/internal/user"
	"MedodsProject/pkg/client/mongodb"
	"MedodsProject/pkg/config"
	"context"
	"log"
)

func main() {
	cfg := config.New()
	ctx := context.TODO()

	mongo, err := mongodb.NewClient(ctx, cfg.MongoDB)
	if err != nil {
		log.Fatal("mongo err: ", err)
	}

	userService := user.NewService(mongo)

	server := http.New(cfg.HTTP, userService)
	log.Fatal(server.Start(ctx))
}
