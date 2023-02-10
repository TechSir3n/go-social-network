package session

import (
	"net/http"
	"social_network/utils/logger"
	"time"
)

func SetCookie(wrt http.ResponseWriter, tokenAccess string) {
	expires := time.Now().AddDate(0, 6, 0) // six month will be avalaible
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

func GetCookie(req *http.Request, key string) *http.Cookie {
	cookie, err := req.Cookie(key)
	if err != nil {
		logger.Fatal("Error occured while reading cookie")
	}

	return cookie
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
