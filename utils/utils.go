package utils

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	templ "html/template"
	"log"
	"net/http"
	"social_network/internal/repository"
	"unicode"
)

func ExecTemplate(w http.ResponseWriter, template string, data interface{}) {
	t, err := templ.ParseFiles(template)

	if err != nil {
		log.Fatal("Failed to read html file")
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return bytes, err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsName(name string) bool {
	sz := len(name)
	fchar := name[0]
	var IsTrue bool = false

	if sz > 6 && unicode.IsUpper(rune(fchar)) {
		IsTrue = true
	} else {
		IsTrue = false
	}

	nm, err := database.GetUser(context.Background())

	if err != nil {
		errors.Wrap(err, "Unable to get data from database")
	}

	for _, iname := range nm {
		if iname.Name == name {
			IsTrue = false
		} else {
			IsTrue = true
		}
	}

	return IsTrue
}

func IsPassword(password string) bool {
	fchar := password[0]
	sz := len(password)

	if sz > 8 && unicode.IsUpper(rune(fchar)) {
		return true
	} else {
		return false
	}
}

func IsEmail(email string) bool {
	login, err := database.GetUser(context.Background())
	var IsTrue bool = false

	if err != nil {
		errors.Wrap(err, "unable to get data from database")
	}

	for _, eml := range login {
		if eml.Email == email {
			IsTrue = false
		}
	}

	return IsTrue
}
