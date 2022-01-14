package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"net/http"
	"nn-blockchain-api/config"
	"nn-blockchain-api/internal/health"
	"nn-blockchain-api/internal/wallet"
	"nn-blockchain-api/pkg/grpc_client"
)

func main() {
	// Init logger
	logger := logrus.New()

	// Init config
	cfg, err := config.Get(".")
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}

	// Set-up gRPC client
	walletClient, err := grpc_client.NewWalletClient(cfg.GRpcHost)
	if err != nil {
		logger.Fatalf("failed to set-up wallet client: %v", err)
	}

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

	// Services
	walletSvc, err := wallet.NewService(walletClient, logger)
	if err != nil {
		logger.Fatalf("failed to create wallet service: %v", err)
	}

	// Handlers
	healthHandler := health.NewHandler()
	walletHandler := wallet.NewHandler(walletSvc)

	router.Route("/api/v1", func(r chi.Router) {
		healthHandler.SetupRoutes(r)
		walletHandler.SetupRoutes(r)
	})

	// Start App
	err = http.ListenAndServe(cfg.PORT, router)
	if err != nil {
		logger.Fatalln("Failed to start HTTP server!")
		return
	}
}
