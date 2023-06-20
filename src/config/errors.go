package config

import (
	"errors"
)

// Define errors
var (
	ErrorServerError    = errors.New("server error")
	ErrorTooManyRequest = errors.New("too many request")
	ErrorUnauthorized   = errors.New("unauthorized")
	ErrorLoginError     = errors.New("incorrect user_id/pwd")
)
