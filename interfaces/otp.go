package interfaces

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"otp/application"
	"otp/domain/entity"
	"otp/sharedinfrastructure/helper"
)

type OtpInterface struct {
	us application.OtpApplication
}

func NewOtp(us application.OtpApplication) OtpInterface {
	return OtpInterface{
		us: us,
	}
}
func (s *OtpInterface) CreateOtp(w http.ResponseWriter, r *http.Request) {
	var otp entity.OtpStruct
	err := decodeJSONBody(w, r, &otp)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			errorResponse := helper.ReturnedError("invalid struct", err.Error())
			log.Println(errorResponse)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&mr)
		} else {
			errorResponse := helper.ReturnedError("invalid struct", err.Error())
			log.Println(errorResponse)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	postOtp, err := s.us.CreateOtp(otp)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(postOtp)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(postOtp)
}

func (s *OtpInterface) ValidateOtp(w http.ResponseWriter, r *http.Request) {
	var otp entity.Auth
	err := decodeJSONBody(w, r, &otp)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			errorResponse := helper.ReturnedError("invalid struct", err.Error())
			log.Println(errorResponse)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&mr)
		} else {
			errorResponse := helper.ReturnedError("invalid struct", err.Error())
			log.Println(errorResponse)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	validateOtp, err := s.us.ValidateOtp(otp)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validateOtp)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(validateOtp)

}
