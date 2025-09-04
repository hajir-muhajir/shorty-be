package domain

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrInactive     = errors.New("link inactive")
	ErrExpired      = errors.New("link expired")
	ErrLimitReached = errors.New("click limit reached")
)
