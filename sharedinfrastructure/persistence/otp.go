package persistence

import (
	"context"
	"errors"
	"fmt"
	"otp/domain/entity"
	"otp/sharedinfrastructure/helper"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OtpInfra struct {
	collection *mongo.Collection
}

func NewOtpInfra(collection *mongo.Collection) *OtpInfra {
	return &OtpInfra{collection}
}

func (r *OtpInfra) CreateOtp(c entity.OtpStruct) (interface{}, error) {

	errorResponse, validateErr := helper.ValidateStruct(c)
	if validateErr != nil {
		return errorResponse, errors.New("")
	}

	c.Reference = uuid.New().String()
	sentAt := time.Now()
	c.SentAt = sentAt
	expiresAt := sentAt.Add(time.Minute * 15)
	c.ExpiresAt = expiresAt
	c.Code = helper.RandomNum()
	fmt.Println(c.Code)
	ref := c.Reference
	_, err := r.collection.InsertOne(context.TODO(), c)
	if err != nil {
		errorResponse := helper.ReturnedError("validation error", "error inserting record into the database")
		return errorResponse, errors.New("")
	}
	return ref, nil
}

func (r *OtpInfra) ValidateOtp(c entity.Auth) (interface{}, error) {
	var s entity.OtpStruct

	filter := bson.M{"user_reference": c.UserReference, "device_reference": c.DeviceReference, "code": c.Code}

	err := r.collection.FindOne(context.TODO(), filter).Decode(&s)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorResponse := helper.ReturnedError("validation error", "invalid code")
			return errorResponse, errors.New("")
		}
	}

	if s.ExpiresAt.Before(time.Now()) {
		errorResponse := helper.ReturnedError("validation error", "expired code")
		return errorResponse, errors.New("")
	}
	return entity.Validity{IsValid: true}, nil
}
