package session

import (
	"net/http"
	"time"
)

func SetCookie(wrt http.ResponseWriter, tokenAccess string) {
	expires := time.Now().AddDate(0, 6, 0)
	cookie := http.Cookie{
		Name:     "Session",
		Value:    tokenAccess,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Expires:  expires,
	}

	http.SetCookie(wrt, &cookie)
}

func ClearCookie(wrt http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "Session",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	}

	http.SetCookie(wrt, &cookie)
}
