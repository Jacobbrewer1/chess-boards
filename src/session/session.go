package session

import (
	"github.com/Jacobbrewer1/chess-boards/src/custom"
	"github.com/google/uuid"
	"net/http"
	"time"
)

var (
	// Sessions this map stores the users Sessions. For larger scale applications, you can use a database or cache for this purpose
	Sessions = custom.Map[string, Session]{}
)

const (
	// UserCookieTimeout Credentials Create a struct that models the structure of a User in the request body
	UserCookieTimeout          = 6
	UserCookieTimoutMultiplier = time.Second
	Timeout                    = UserCookieTimeout * UserCookieTimoutMultiplier

	test = UserCookieTimoutMultiplier | time.Hour
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(CookieTokenName)
	if err != nil {
		if err == http.ErrNoCookie {
			// If there are no cookies then return and handle redirects in calling method
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	oldToken := c.Value

	userSession, exists := Sessions[oldToken]
	if !exists {
		// If there are no cookies then return and handle redirects in calling method
		return
	}

	if userSession.IsExpired() {
		delete(Sessions, oldToken)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	newToken := uuid.NewString()
	expiresAt := time.Now().UTC().Add(Timeout)

	Sessions[newToken] = NewSession(userSession.User, expiresAt).GetSession()

	delete(Sessions, oldToken)

	http.SetCookie(w, &http.Cookie{
		Name:    CookieTokenName,
		Value:   newToken,
		Expires: expiresAt,
		Path:    "/",
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	go clearExpiredSessions()

	c, err := r.Cookie(CookieTokenName)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	sessionToken := c.Value

	delete(Sessions, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    CookieTokenName,
		Value:   "",
		Expires: time.Now(),
		Path:    "/",
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
