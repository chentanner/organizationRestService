package services

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

//ServiceValidatorInterface interface for validator
type ServiceValidatorInterface interface {
	ValidateStruct(structToValidate interface{}) error
}

// ServiceValidator type containing ServiceValidator
type ServiceValidator struct {
	validator *validator.Validate
}

// NewServiceValidator creates a new instance
func NewServiceValidator() *ServiceValidator {
	serviceValidator := new(ServiceValidator)
	serviceValidator.validator = validator.New()
	return serviceValidator
}

func (serviceValidator *ServiceValidator) ValidateStruct(structToValidate interface{}) error {
	err := serviceValidator.validator.Struct(structToValidate)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}

		fmt.Println()
		fmt.Println(err.Error())

		return err
	}

	return nil
}
