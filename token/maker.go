package token

import "time"

type Maker interface {
	// トークンを作る
	CreateToken(username string, duratio time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
