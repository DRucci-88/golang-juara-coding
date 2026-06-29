package domain

import "time"

type BlackListedToken struct {
	ID          uint
	TokenString string
	ExpireAt    time.Time
}

type BlackListedTokenRepository interface {
	Create(token string, expireAt time.Time) error
	IsTokenBlacklisted(token string) (bool, error)
	DeleteExpired() (int64, error)
}
