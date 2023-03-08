package session

import (
	"errors"
	"github.com/Jacobbrewer1/chess-boards/src/entities"
	"log"
	"net/http"
)

const CookieTokenName string = "racketHut_session_token"

func GetUserFromCookies(r *http.Request) (entities.User, error) {
	t, err := r.Cookie(CookieTokenName)
	if err != nil {
		log.Println(err)
		return entities.User{}, err
	}

	if _, ok := Sessions[t.Value]; !ok {
		return entities.User{}, errors.New("invalid session token")
	}

	return Sessions[t.Value].User, nil
}

func UpdateUserSession(r *http.Request, newUser entities.User) error {
	t, err := r.Cookie(CookieTokenName)
	if err != nil {
		log.Println(err)
		return err
	}

	Sessions[t.Value] = Session{
		User:   newUser,
		Expiry: Sessions[t.Value].Expiry,
	}

	return nil
}

func clearExpiredSessions() {
	var toRemove = make([]string, len(Sessions))

	for key, val := range Sessions {
		if !val.IsExpired() {
			continue
		}

		toRemove = append(toRemove, key)
	}

	if toRemove != nil && len(toRemove) > 0 {
		for _, key := range toRemove {
			delete(Sessions, key)
		}
	}
}
