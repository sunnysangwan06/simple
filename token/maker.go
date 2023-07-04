package token

import (
	"time"
)

// Maker is a interface that creates and verifies tokens.
type Maker interface {
	// CreateToken creates a new token for a specific username and duration.
	CreateToken(username string, duration time.Duration) (string,  error)

	// VerifyToken checks if the token is valid or not.
	VerifyToken(token string) (*Payload, error)
}