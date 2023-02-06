package utils

import (
	"context"
	"github.com/pkg/errors"
	templ "html/template"
	"log"
	"net/http"

	"social_network/internal/api/v1/models"
	"social_network/internal/repository/database/postgresql"
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

func IsName(name string) bool {
	sz := len(name)
	fchar := name[0]
	var IsTrue bool = false

	if sz > 6 && unicode.IsUpper(rune(fchar)) {
		IsTrue = true
	} else {
		IsTrue = false
	}

	return IsTrue
}

func IsPassword(password string) bool {
	fchar := password[0]
	sz := len(password)

	if sz > 7 && unicode.IsUpper(rune(fchar)) {
		return true
	} else {
		return false
	}
}

func IsEmail(email string) bool {
	login, err := database.GetUser(context.Background())
	var IsTrue bool = true

	if err != nil {
		errors.Wrap(err, "unable to get data from database")
		IsTrue = false
	}

	for _, eml := range login {
		if eml.Email == email {
			IsTrue = false
		}
	}

	return IsTrue
}

func Validate(user *models.User) error {
	if user.Name == "" {
		return errors.New("name is empty")
	} else if user.Email == "" {
		return errors.New("email is empty")
	} else if user.Password == "" {
		return errors.New("passowrd is empty")
	} else if user.ConfirmPassword == "" {
		return errors.New("confirm Password is empty")
	} else {
		return nil
	}
}
