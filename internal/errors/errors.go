package errors

import "errors"

var (
	ErrDBNotInitialized = errors.New("database not initialized")
)
