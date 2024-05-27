package smpp

import "otp-smpp-go/internal/config"

type SMPPClient interface {
	SendSMS(cfg *config.Config, dest, text string) error
}
