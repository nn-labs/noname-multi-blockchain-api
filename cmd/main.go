package main

import (
	"net/http"
	"nn-blockchain-api/config"
	"nn-blockchain-api/internal/bitcoin"
	"nn-blockchain-api/internal/health"
	"nn-blockchain-api/internal/wallet"
	"nn-blockchain-api/pkg/grpc_client"
	bitcoin_rpc "nn-blockchain-api/pkg/rpc/bitcoin"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

func main() {
	// Init logger
	logger := logrus.New()

	// Init config
	cfg, err := config.Get()
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}

	// Set-up gRPC client
	walletClient, err := grpc_client.NewWalletClient(cfg.GRps.GRpcHost)
	if err != nil {
		logger.Fatalf("failed to set-up wallet client: %v", err)
	}

	// Rpc clients
	btcRpcClient, err := bitcoin_rpc.NewBtcClient(cfg.BtcRpc.BtcRpcEndpointTest, cfg.BtcRpc.BtcRpcEndpointMain, cfg.BtcRpc.BtcRpcUser, cfg.BtcRpc.BtcRpcPassword)
	if err != nil {
		logger.Fatalf("failed to set-up btc rpc client: %v", err)
	}

	// Rpc services
	bitcoinRpcTxSvc, err := bitcoin_rpc.NewTransactionService(btcRpcClient)
	if err != nil {
		logger.Fatalf("failed to create bitcoin tx service: %v", err)
	}
	bitcoinRpcWalletSvc, err := bitcoin_rpc.NewWalletService(btcRpcClient)
	if err != nil {
		logger.Fatalf("failed to create bitcoin wallet service: %v", err)
	}
	bitcoinRpcHealthSvc, err := bitcoin_rpc.NewHealthService(btcRpcClient)
	if err != nil {
		logger.Fatalf("failed to create bitcoin health service: %v", err)
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

	bitcoinSvc, err := bitcoin.NewService(logger, bitcoinRpcTxSvc, bitcoinRpcWalletSvc, bitcoinRpcHealthSvc)
	if err != nil {
		logger.Fatalf("failed to create bitcoin service: %v", err)
	}

	// Handlers
	healthHandler := health.NewHandler()
	walletHandler := wallet.NewHandler(walletSvc)
	bitcoinHandler := bitcoin.NewHandler(bitcoinSvc)

	router.Route("/api/v1", func(r chi.Router) {
		healthHandler.SetupRoutes(r)
		walletHandler.SetupRoutes(r)
	})

	router.Route("/api/v1/bitcoin", func(r chi.Router) {
		bitcoinHandler.SetupRoutes(r)
	})

	// Start App
	err = http.ListenAndServe(cfg.PORT, router)
	if err != nil {
		logger.Fatalln("Failed to start HTTP server!")
		return
	}
}
