package token

import "time"

// Maker is an interface for managing token
type Maker interface {
	// CreateToken func create token for specific name and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// VerifyToken check if token is invalid or not
	VerifyToken(token string) (*Payload, error)
}
