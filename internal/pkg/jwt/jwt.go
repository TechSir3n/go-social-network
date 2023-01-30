package jwt

import (
	"fmt"
	"os"
	_ "social_network/internal/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTSecret = os.Getenv("SECKET_KEY")

func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(JWTSecret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

//func IsValidToken(token string) bool{}
