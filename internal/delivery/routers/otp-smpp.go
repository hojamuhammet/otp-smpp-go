package routers

import (
	"net/http"
	"otp-smpp-go/internal/delivery/handlers"
	"otp-smpp-go/internal/service"
	"otp-smpp-go/pkg/lib/logger"

	"github.com/go-chi/chi/v5"
)

func SetupOTPRoutes(otpService service.OTPService, logger *logger.Loggers) http.Handler {
	otpRouter := chi.NewRouter()
	otpHandler := handlers.NewOTPHandler(otpService)

	otpRouter.Post("/sendOTP", otpHandler.GenerateAndSaveOTPHandler)

	return otpRouter
}
