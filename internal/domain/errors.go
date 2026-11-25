package domain

import "errors"

var (
	ErrNotFound   = errors.New("Not Found")
	ErrEmailTaken = errors.New("email already registred")
)
