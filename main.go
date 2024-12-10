package main

import (
	"bank-simulation/api"
	"bank-simulation/config"
	"bank-simulation/handlers"
	"bank-simulation/repository"
	"log"
	"net/http"
)

func main() {
	log.Println("Server running on :8000")

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Could not connect to the daabase : %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	accountRepo := repository.NewAccountRepository(db)
	accountHandler := handlers.NewAccountHandler(accountRepo)

	router := api.SetupRouter(userHandler, accountHandler)

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
}
