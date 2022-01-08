package helper

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/google/uuid"
)

func (r ErrorResponse) ErrorMessage(e ErrorBody, code string, message interface{}, source string) ErrorResponse {
	r.ErrorReference = uuid.New()
	r.TimeStamp = time.Now().Format(time.RFC3339)
	e.Code = code
	e.Message = message
	e.Source = source
	r.Errors = append(r.Errors, e)
	return r
}

func ReturnedError(code string, message interface{}) interface{} {
	var r ErrorResponse
	var e ErrorBody
	r.ErrorMessage(e, code, message, "otp-service")
	return r
}   
func (r SuccessBody) SuccessMessage(code, message string, reference string) SuccessBody {
	r.Code = code
	r.Message = message
	r.Reference = reference
	return r
}
func (r AllSuccessBody) AllSuccessMessage(code, message string, service interface{}) AllSuccessBody {
	r.Code = code
	r.Message = message
	r.Service = service
	return r
}
func validate() (*validator.Validate, ut.Translator) {
	validate := validator.New()
	readableErr := en.New()
	uni := ut.New(readableErr, readableErr)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatal(err)
	}
	return validate, trans
}

func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

func ValidateStruct(data interface{}) (interface{}, error) {
	//validation here
	validate, trans := validate()
	err := validate.Struct(data)
	errorValue := translateError(err, trans)
	//return errors here
	return returnError(errorValue), err

}

func returnError(errorArray []error) ErrorResponse {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()

	for _, value := range errorArray {
		body := ErrorBody{Code: "validation error", Source: "otp-service", Message: value.Error()}
		errorResponse.Errors = append(errorResponse.Errors, body)
	}
	return errorResponse
}

func RandomNum() int {
	rand.Seed(time.Now().UnixNano())
	min := 100001
	max := 999999
	x := rand.Intn(max-min+1) + min
	return x
}
