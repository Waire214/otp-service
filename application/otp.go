package application

import (
	"otp/domain/entity"
	"otp/domain/repository"
)

type OtpApp struct {
	theOtp repository.OtpRepository
}

var _OtpApplication = &OtpApp{}

type OtpApplication interface {
	CreateOtp(entity.OtpStruct) (interface{}, error)
	ValidateOtp(entity.Auth) (interface{}, error)
}

func (u *OtpApp) CreateOtp(c entity.OtpStruct) (interface{}, error) {
	return u.theOtp.CreateOtp(c)
}

func (u *OtpApp) ValidateOtp(c entity.Auth) (interface{}, error) {
	return u.theOtp.ValidateOtp(c)
}