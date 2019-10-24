package appErrors

type TransactionErrors struct{
	transactionErrors map[string]TransactionError
}

type TransactionError struct{
	Code string
	Message string
}

func NewTransactionError(code string, message string)TransactionError{
	transactionError := new(TransactionError)
	transactionError.Code = code
	transactionError.Message = message
	return *transactionError
}


var MISSING_ORGANIZATION_SHORT_NAME TransactionError = NewTransactionError(	"E11000", "Missing Organization short name.")
var MISSING_ORGANIZATION_LEGAL_NAME TransactionError = NewTransactionError(	"E11001", "Missing Organization legal name.")
var DUPLICATE_ORG_SHORT_NAME TransactionError = NewTransactionError(	"E11002", "Duplicate organization short names are not allowed. This short name is currently in use. Please enter another.")
var DUPLICATE_ORG_LEGAL_NAME TransactionError = NewTransactionError(	"E11003", "Duplicate organization legal names are not allowed. This legal name is currently in use. Please enter another.")
var INVALID_ORG_SHORT_NAME TransactionError = NewTransactionError(	"E11004", "The organization short name is not valid.")
