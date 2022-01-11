package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"nn-blockchain-api/config"
	"nn-blockchain-api/internal/health"
)

func main() {
	// Init config
	cfg, err := config.Get(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Println(cfg)

	// Set-up Route
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"OPTIONS", "GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
		ExposedHeaders:   []string{"Content-Type", "JWT-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	//router.Use(middleware.BasicAuth("authentication", map[string]string{cfg.User: cfg.Password}))

	// Handlers
	healthHandler := health.NewHandler()
	healthHandler.SetupRoutes(router)

	// Start App
	err = http.ListenAndServe(cfg.PORT, router)
	if err != nil {
		panic("ERROR: Something wrong with start app!")
	}
}
