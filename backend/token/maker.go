package token

import "time"

type Maker interface {
	//Create a new token
	CreateToken(username string, duration time.Duration) (string, error)

	//Verify the token
	VerifyToken(token string) (*Payload, error)
}
