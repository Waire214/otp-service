package entity

import "time"

type OtpStruct struct {
	Reference       string    `bson:"reference" json:"reference"`
	Code            int       `bson:"code" json:"code"`
	Strict          bool      `bson:"strict" json:"strict"`
	SentAt          time.Time `bson:"sent_at" json:"sent_at"`
	ExpiresAt       time.Time `bson:"expires_at" json:"expires_at"`
	DeviceReference string    `bson:"device_reference" json:"device_reference"`
	UserReference   string    `bson:"user_reference" json:"user_reference"`
}

type Validity struct {
	IsValid bool `bson:"is_valid" json:"is_valid"`
}

type Auth struct {
	Code            int    `bson:"code" json:"code"`
	DeviceReference string `bson:"device_reference" json:"device_reference"`
	UserReference   string `bson:"user_reference" json:"user_reference"`
}
