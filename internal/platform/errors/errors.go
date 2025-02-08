package errors

import "errors"

// Common errors
var (
	ErrConnecting  = errors.New("error connecting to MongoDB")
	ErrPinging     = errors.New("error pinging MongoDB")
	ErrNoDocuments = errors.New("no documents found") // Example of a common error
)

// Define common error types
var (
	ErrNotFound            = errors.New("not found")
	ErrInvalidInput        = errors.New("invalid input")
	ErrInternalServerError = errors.New("internal server error")
	// Add more specific error types as needed
)
