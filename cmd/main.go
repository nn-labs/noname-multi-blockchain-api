package main

import (
	"errors"
	"go.uber.org/zap"
	"log"
	"net/http"
	"nn-blockchain-api/config"
	"nn-blockchain-api/internal/bitcoin"
	"nn-blockchain-api/internal/ethereum"
	"nn-blockchain-api/internal/health"
	"nn-blockchain-api/internal/wallet"
	"nn-blockchain-api/pkg/grpc_client"
	"nn-blockchain-api/pkg/logger"
	bitcoin_rpc "nn-blockchain-api/pkg/rpc/bitcoin"
	ethereum_rpc "nn-blockchain-api/pkg/rpc/ethereum"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	// Init config
	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Init logger
	newLogger, err := logger.NewLogger(cfg.AppEnv)
	if err != nil {
		log.Fatalf("can't create logger: %v", err)
	}

	zapLogger, err := newLogger.SetupZapLogger()
	if err != nil {
		log.Fatalf("can't setup zap logger: %v", err)
	}
	defer func(zapLogger *zap.SugaredLogger) {
		err := zapLogger.Sync()
		if err != nil && !errors.Is(err, syscall.ENOTTY) {
			log.Fatalf("can't setup zap logger: %v", err)
		}
	}(zapLogger)

	// Set-up gRPC client
	walletClient, err := grpc_client.NewWalletClient(cfg.GRps.GRpcHost)
	if err != nil {
		zapLogger.Fatalf("failed to set-up wallet client: %v", err)
	}

	// Rpc clients
	bitcoinRpcClient, err := bitcoin_rpc.NewClient(cfg.BtcRpc.BtcRpcEndpointTest, cfg.BtcRpc.BtcRpcEndpointMain, cfg.BtcRpc.BtcRpcUser, cfg.BtcRpc.BtcRpcPassword)
	if err != nil {
		zapLogger.Fatalf("failed to set-up btc rpc client: %v", err)
	}

	ethereumRpcClient, err := ethereum_rpc.NewClient(cfg.EthRpc.EthRpcEndpointTest, cfg.EthRpc.EthRpcEndpointMain)
	if err != nil {
		zapLogger.Fatalf("failed to set-up btc rpc client: %v", err)
	}

	// Rpc services
	bitcoinRpcService, err := bitcoin_rpc.NewService(bitcoinRpcClient)
	if err != nil {
		zapLogger.Fatalf("failed to create bitcoin service: %v", err)
	}

	ethereumRpcService, err := ethereum_rpc.NewService(ethereumRpcClient)
	if err != nil {
		zapLogger.Fatalf("failed to create bitcoin service: %v", err)
	}

	// Services
	walletService, err := wallet.NewService(walletClient, zapLogger)
	if err != nil {
		zapLogger.Fatalf("failed to create wallet service: %v", err)
	}

	bitcoinService, err := bitcoin.NewService(bitcoinRpcService, zapLogger)
	if err != nil {
		zapLogger.Fatalf("failed to create bitcoin service: %v", err)
	}

	ethereumService, err := ethereum.NewService(ethereumRpcService, zapLogger)
	if err != nil {
		zapLogger.Fatalf("failed to create bitcoin service: %v", err)
	}

	// Handlers
	healthHandler := health.NewHandler()

	walletHandler, err := wallet.NewHandler(walletService)
	if err != nil {
		zapLogger.Fatalf("failed to create wallet handler: %v", err)
	}

	bitcoinHandler, err := bitcoin.NewHandler(bitcoinService)
	if err != nil {
		zapLogger.Fatalf("failed to create bitcoin handler: %v", err)
	}

	ethereumHandler, err := ethereum.NewHandler(ethereumService)
	if err != nil {
		zapLogger.Fatalf("failed to create bitcoin handler: %v", err)
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

	router.Route("/api/v1", func(r chi.Router) {
		healthHandler.SetupRoutes(r)
		walletHandler.SetupRoutes(r)
	})

	router.Route("/api/v1/bitcoin", func(r chi.Router) {
		bitcoinHandler.SetupRoutes(r)
	})

	router.Route("/api/v1/ethereum", func(r chi.Router) {
		ethereumHandler.SetupRoutes(r)
	})

	// Start App
	err = http.ListenAndServe(cfg.PORT, router)
	if err != nil {
		zapLogger.Fatal("Failed to start HTTP server!")
		return
	}
}
