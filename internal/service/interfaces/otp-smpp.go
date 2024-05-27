package service

import (
	"otp-smpp-go/internal/config"
)

type OTPService interface {
	SendOTP(cfg config.Config, phoneNumber string) error
}
