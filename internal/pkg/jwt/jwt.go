package jwt

import (
	"log"
	"os"
	"social_network/internal/api/v1/models"
	_ "social_network/internal/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct { // it's custom paylod in jwt token
	Email    string `json:"Email"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECKET_KEY")))

	if err != nil {
		log.Println(err.Error())
		return "", nil
	}

	return tokenString, nil
}

func IsValidToken(tokenStr string) (*jwt.Token, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECKET_KEY")), nil
	})

	if err != nil {
		log.Println(err, " :Invalid Authorization")
		return nil, err
	}

	if !token.Valid {
		log.Println(err, " :Invalid Token")
		return nil, err
	}

	return token, err
}
