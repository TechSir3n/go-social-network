package session

import (
	"net/http"
	"time"
)

func SetCookie(wrt http.ResponseWriter, token string) {
	expires := time.Now().AddDate(0, 6, 0)
	cookie := http.Cookie{
		Name:     "Access_Token",
		Value:    token,
		
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Expires:  expires,
	}

	http.SetCookie(wrt, &cookie)
}

func ClearCookie(wrt http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "Access_Token",
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	}

	http.SetCookie(wrt, &cookie)

}
