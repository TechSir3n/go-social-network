package v2

import (
	"net/http"
	"time"
)

func SetCookie(wrt http.ResponseWriter, token string) {
	expires := time.Now().AddDate(1, 0, 0)
	cookie := http.Cookie{
		Name:     "Access_Token",
		Value:    token,
		HttpOnly: true,
		Path:     "/login",
		Expires:  expires,
	}

	http.SetCookie(wrt, &cookie)
}

func ClearCookie(wrt http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "Access_Token",
		HttpOnly: true,
		Path:     "/login",
		MaxAge:   -1,
	}

	http.SetCookie(wrt, &cookie)

}
