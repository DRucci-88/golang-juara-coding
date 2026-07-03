package model

import (
	"errors"
)

var (
	ErrAuthUnauthorized             = errors.New("Email or Password Wrong")
	ErrAuthWrongAuthorizationHeader = errors.New("Wrong Header Authorization")
)
