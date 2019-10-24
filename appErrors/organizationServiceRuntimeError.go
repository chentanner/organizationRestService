package appErrors

import (
	"encoding/json"
)

// OrganizationRestServiceRuntimeError interface for application errors
type OrganizationRestServiceRuntimeError interface {
	Error() string
	ErrorCode() string
	ErrorMessage() string
	ErrorParams() []string
	TransactionErrorMessage() string
	TransactionError() TransactionError
}

// Error returns InvalidValidationError message
type organizationRestServiceRuntimeError struct {
	errorParams      []string
	transactionError TransactionError
}

// NewRuntimeError creates a new error with a transactionerror
func NewRuntimeError(transactionError TransactionError) OrganizationRestServiceRuntimeError {
	err := new(organizationRestServiceRuntimeError)
	err.transactionError = transactionError
	return err
}

// NewRuntimeErrorWithParams creates a new error with a transactionerror and an array of params
func NewRuntimeErrorWithParams(transactionError TransactionError, params []string) OrganizationRestServiceRuntimeError {
	err := new(organizationRestServiceRuntimeError)
	err.transactionError = transactionError
	err.errorParams = params
	return err
}

// Error implementation for the golang errors interface
func (e *organizationRestServiceRuntimeError) Error() string {
	return e.ErrorMessage()
}

func (e *organizationRestServiceRuntimeError) ErrorCode() string {
	return e.transactionError.Code
}

func (e *organizationRestServiceRuntimeError) ErrorMessage() string {
	return e.transactionError.Message
}

func (e *organizationRestServiceRuntimeError) ErrorParams() []string {
	return e.errorParams
}

func (e *organizationRestServiceRuntimeError) TransactionErrorMessage() string {
	return e.transactionError.Message
}

func (e *organizationRestServiceRuntimeError) TransactionError() TransactionError {
	return e.transactionError
}

func (e *organizationRestServiceRuntimeError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    e.ErrorCode(),
		Message: e.ErrorMessage(),
	})
}
