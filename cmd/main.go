package main

import (
	"log/slog"
	"net/http"
	"os"
	"otp-smpp-go/internal/config"
	"otp-smpp-go/internal/delivery/routers"
	smpp "otp-smpp-go/internal/infrastructure"
	"otp-smpp-go/internal/service"
	"otp-smpp-go/pkg/lib/logger"
)

func main() {
	cfg := config.LoadConfig()

	logger, err := logger.SetupLogger(cfg.Env)
	if err != nil {
		slog.Error("failed to set up logger: %v", err)
		os.Exit(1)
	}

	logger.InfoLogger.Info("Server is up and running")
	slog.Info("Server is up and running")

	smppClient, err := smpp.NewSMPPClient(cfg)
	if err != nil {
		logger.ErrorLogger.Error("failed to initialize SMPP client: %v", err)
		os.Exit(1)
	}

	otpService := service.NewOTPService(logger, cfg, smppClient)

	r := routers.SetupOTPRoutes(otpService, logger)

	err = http.ListenAndServe(cfg.HTTPServer.Address, r)
	if err != nil {
		logger.ErrorLogger.Error("Server failed to start:", err)
	}
}
