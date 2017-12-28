package models

import "errors"

func falseAndError(msg string) (bool, error) {
	return false, errors.New(msg)
}

func raiseIfError(err error) {
	if err != nil {
		panic(err)
	}
}
