package jwt

import (
	"os"
	"time"

	"social_network/internal/api/v1/models"
	"social_network/utils/logger"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

// Generate JWT(json web token)
func GenerateJWT(user models.User) (*PayloadJWT, error) {
	payload := &PayloadJWT{}
	var err error

	payload.ExpiresAccess = time.Now().Add(30 * time.Minute).Unix()
	payload.AccessUID = uuid.NewV4().String()

	access_token := jwt.MapClaims{}
	access_token["authorized"] = true
	access_token["user_email"] = user.Email
	access_token["access_uuid"] = payload.AccessUID
	access_token["exp"] = payload.ExpiresRefresh

	strAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, access_token)
	payload.AccessToken, err = strAccess.SignedString([]byte(os.Getenv("ACCESS_TOKEN")))

	if err != nil {
		logger.Error(err, " :[ERROR] ACCESS TOKEN")
		return &PayloadJWT{}, nil
	}

	payload.ExpiresRefresh = time.Now().Add(30 * time.Minute).Unix()
	payload.RefreshUID = uuid.NewV4().String()

	refresh_token := jwt.MapClaims{}
	refresh_token["user_email"] = user.Email
	refresh_token["refresh_uuid"] = payload.RefreshUID
	refresh_token["exp"] = payload.ExpiresRefresh

	strRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, access_token)
	payload.RefreshToken, err = strRefresh.SignedString([]byte(os.Getenv("REFRESH_TOKEN")))
	if err != nil {
		logger.Error(err, " :[ERROR] REFRESH TOKEN")
		return &PayloadJWT{}, nil
	}

	return payload, nil
}

func ParseJWT(tokenStr string) (*jwt.Token, error) {
	payload := &PayloadJWT{}
	token, err := jwt.ParseWithClaims(tokenStr, payload, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_TOKEN")), nil
	})

	if err != nil {
		logger.Error(err, " :Invalid Authorization")
		return nil, err
	}

	if !token.Valid {
		logger.Error(err, " :Invalid Token")
		return nil, err
	}

	return token, err
}
