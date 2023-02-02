package jwt

import "github.com/dgrijalva/jwt-go"

type PayloadJWT struct {	 // it's custom paylod in jwt token
	Email          string `json:"Email"`
	Name           string `json:"Name"`
	AccessToken    string `json:"AccessToken"`
	RefreshToken   string `json:"RefreshToken"`
	AccessUID      string `json:"AccessUID"`
	RefreshUID     string `json:"RefreshUID"`
	ExpiresAccess  int64  `json:"ExpiresAccess"`
	ExpiresRefresh int64  `json:"ExpiresRefresh"`
	jwt.StandardClaims
}
