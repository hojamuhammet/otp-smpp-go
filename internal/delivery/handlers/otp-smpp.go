package handlers

import (
	"encoding/json"
	"net/http"
	"otp-smpp-go/internal/service"
	"otp-smpp-go/pkg/lib/utils"
)

type OTPHandler struct {
	Service service.OTPService
}

func NewOTPHandler(s service.OTPService) *OTPHandler {
	return &OTPHandler{Service: s}
}

func (h *OTPHandler) GenerateAndSaveOTPHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		PhoneNumber string `json:"phone_number"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.RespondWithErrorJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if request.PhoneNumber == "" {
		utils.RespondWithErrorJSON(w, http.StatusBadRequest, "Phone number is required")
		return
	}

	err = h.Service.SendOTP(request.PhoneNumber)
	if err != nil {
		utils.RespondWithErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
