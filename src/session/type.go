package session

import (
	"github.com/Jacobbrewer1/chess-boards/src/custom"
	"github.com/Jacobbrewer1/chess-boards/src/entities"
	"time"
)

type Session struct {
	Key    string          `json:"key"`
	User   entities.User   `json:"user"`
	Expiry custom.Datetime `json:"expiry"`
}

func NewSession(user entities.User, expiry time.Time) *Session {
	return &Session{
		User:   user,
		Expiry: (custom.Datetime)(expiry),
	}
}

// IsExpired we'll use this method later to determine if the session has expired
func (s Session) IsExpired() bool {
	return s.Expiry.TimeValue().Before(time.Now().UTC())
}

func (s Session) GetSession() Session {
	return s
}
