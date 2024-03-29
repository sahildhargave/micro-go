package token

import "time"

//Maker is an interface for managing token

type Maker interface {
	// CreateToken create a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)
	// VerifyToken check if the token is valid and return else not
	VerifyToken(Token string) (*Payload, error)
}
