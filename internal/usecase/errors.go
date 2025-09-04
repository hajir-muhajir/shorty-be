package usecase

import "errors"

var(
	ErrUnauthorized = errors.New("unauthorized")
	ErrConflict = errors.New("conflict")
)