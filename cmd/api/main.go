package main

import (
	"cdn-api/internal/handlers"
	"cdn-api/internal/services"
	"log"
	"net/http"
)

func main() {
	userService := &services.UserService{}

	userHandler := &handlers.UserHandler{
		Service: userService,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", userHandler.HandleCreate)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start :( %v)", err)
	}
}
