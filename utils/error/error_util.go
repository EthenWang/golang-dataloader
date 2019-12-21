package error

import "errors"

type statusCode uint

const (
	InvalidDataType statusCode = iota
	DataIsNil
	LanguageNotExist
	TranslationExist
)

type dataLoaderError struct {
	errorCode statusCode
	error
}

func NewDataLoaderError(code statusCode, params ...string) error {
	var errStr string
	switch code {
	case TranslationExist:
		errStr = "Translation exits"
	}
	return &dataLoaderError{code, errors.New(errStr)}
}
