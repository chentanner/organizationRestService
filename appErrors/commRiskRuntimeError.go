package appErrors

import (
	"encoding/json"
)

type organizationRestServiceRuntimeError interface {
	Error() string
	ErrorCode() string
	ErrorMessage() string
	ErrorParams() []string
	TransactionErrorMessage() string
	TransactionError() TransactionError
}

// Error returns InvalidValidationError message
type organizationRestServiceRuntimeError struct {
	errorParams []string
	transactionError TransactionError
}

// NeworganizationRestServiceRuntimeError creates a new error with a transactionerror
func NewRuntimeError(transactionError TransactionError) organizationRestServiceRuntimeError{
	err := new(organizationRestServiceRuntimeError)
	err.transactionError = transactionError
	return err
}

// NeworganizationRestServiceRuntimeError creates a new error with a transactionerror and an array of params
func NewRuntimeErrorWithParams(transactionError TransactionError, params []string) organizationRestServiceRuntimeError{
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
        Code     string   `json:"code"`
        Message  string   `json:"message"`
    }{
        Code: e.ErrorCode(),
        Message:  e.ErrorMessage(),
    })
}