package main

import (
	"github.com/satoshi1975/smartChat/common/auth"
	"github.com/satoshi1975/smartChat/services/auth-service/config"
	"github.com/satoshi1975/smartChat/services/auth-service/internal/db"
	"github.com/satoshi1975/smartChat/services/auth-service/internal/repository"
	"github.com/satoshi1975/smartChat/services/auth-service/internal/services"
	"github.com/satoshi1975/smartChat/services/auth-service/pkg/handlers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db, err := db.NewPostgresDB(&cfg.DB)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	jwtService := auth.NewJWTService(cfg.JWT.SecretKey)
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo, jwtService)
	userHandler := handlers.NewUserHandler(userService)

	router := httprouter.New()
	router.POST("/users", userHandler.CreateUser)
	router.POST("/login", userHandler.Login)

	log.Fatal(http.ListenAndServe(":8080", router))
}
