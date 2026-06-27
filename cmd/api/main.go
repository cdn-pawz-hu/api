package main

import (
	"cdn-api/internal/config"
	"cdn-api/internal/handlers"
	"cdn-api/internal/models"
	"cdn-api/internal/repositories"
	"cdn-api/internal/services"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	cfg config.PostgresConfig
}

func main() {
	config := config.Load()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		config.Postgres.Host,
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.DBName,
		config.Postgres.Port,
		config.Postgres.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	validator := validator.New()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService, validator)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", userHandler.Register)

	log.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
