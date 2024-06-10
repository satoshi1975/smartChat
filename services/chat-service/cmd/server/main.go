package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/satoshi1975/smartChat/common/auth"
	"github.com/satoshi1975/smartChat/services/chat-service/config"
	_ "github.com/satoshi1975/smartChat/services/chat-service/docs"
	"github.com/satoshi1975/smartChat/services/chat-service/internal/db"
	"github.com/satoshi1975/smartChat/services/chat-service/internal/middleware"
	"github.com/satoshi1975/smartChat/services/chat-service/internal/repository"
	"github.com/satoshi1975/smartChat/services/chat-service/internal/services"
	ws "github.com/satoshi1975/smartChat/services/chat-service/internal/websocket"
	"github.com/satoshi1975/smartChat/services/chat-service/pkg/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title SmartChat Profile Service API
// @version 1.0
// @description This is a sample profile service for SmartChat.
// @host localhost:8081
// @BasePath /
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
	jwtMiddleware := middleware.NewJWTMiddleware(jwtService)

	profileRepo := repository.NewProfileRepository(db)
	profileService := services.NewProfileService(profileRepo)
	profileHandler := handlers.NewProfileHandler(profileService)

	hub := ws.NewHub()
	go hub.Run()

	wsHandler := handlers.NewWebSocketHandler(hub)

	router := httprouter.New()
	router.POST("/profiles", jwtMiddleware.RequireAuth(profileHandler.CreateProfile))
	router.GET("/profiles/:id", jwtMiddleware.RequireAuth(profileHandler.GetProfile))
	router.PUT("/profiles/:id", jwtMiddleware.RequireAuth(profileHandler.UpdateProfile))
	router.DELETE("/profiles/:id", jwtMiddleware.RequireAuth(profileHandler.DeleteProfile))
	router.POST("/profiles/:id/friends", jwtMiddleware.RequireAuth(profileHandler.AddFriend))
	router.POST("/profiles/:id/block", jwtMiddleware.RequireAuth(profileHandler.BlockUser))

	router.GET("/ws", wsHandler.ServeWebSocket)

	// Swagger documentation route
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8081", router))
}
