package token

import "time"

type Maker interface {
	//Create a new token
	CreateToken(username string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
