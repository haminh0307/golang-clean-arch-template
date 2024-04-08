package domain

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrWrongCredentials = errors.New("wrong credentials")
)
