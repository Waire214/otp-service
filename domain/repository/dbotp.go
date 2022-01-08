package repository

import "otp/domain/entity"

type OtpRepository interface {
	CreateOtp(entity.OtpStruct) (interface{}, error)
	ValidateOtp(entity.Auth) (interface{}, error)
}