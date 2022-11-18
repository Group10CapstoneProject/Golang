package myerrors

import "errors"

var (
	ErrEmailAlredyExist = errors.New("email is used")
)
