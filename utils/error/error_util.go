package error

import "errors"

type statusCode uint

const (
	LanguageNotExist statusCode = 1
	TranslationExist statusCode = 2
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
