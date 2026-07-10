package domain

import "errors"

var (
	ErrNotFound         = errors.New("resource not found")
	ErrAlreadyTaken     = errors.New("resource already taken")
	ErrInvalidOperation = errors.New("invalid operation")
)
