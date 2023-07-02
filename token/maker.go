package token

import "time"

type Maker interface {
	// トークンを作る
	CreateToken(username string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
