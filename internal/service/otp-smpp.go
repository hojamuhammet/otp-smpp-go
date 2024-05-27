package service

import (
	"fmt"
	"math/rand"
	"otp-smpp-go/internal/config"
	smpp "otp-smpp-go/internal/infrastructure/interfaces"
	"otp-smpp-go/pkg/lib/logger"
	"time"
)

type OTPService struct {
	logger     *logger.Loggers
	cfg        *config.Config
	smppClient smpp.SMPPClient
}

func NewOTPService(logger *logger.Loggers, cfg *config.Config, smppClient smpp.SMPPClient) OTPService {
	return OTPService{logger: logger, cfg: cfg, smppClient: smppClient}
}

func GenerateOTP() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := fmt.Sprintf("%06d", r.Intn(1000000))
	return otp
}

func (s *OTPService) SendOTP(phoneNumber string) error {
	otp := GenerateOTP()

	err := s.smppClient.SendSMS(s.cfg, phoneNumber, otp)
	if err != nil {
		s.logger.ErrorLogger.Error("Error sending OTP via SMS: %v", err)
		return err
	}

	return nil
}
