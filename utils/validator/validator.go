package validator

import (
	"github.com/go-playground/validator/v10"
	"social_network/internal/api/v1/models"
	"social_network/utils/logger"
)

var Validator = validator.New()

func CheckIsValid(user *models.User) error {
	err := Validator.Struct(user)
	if err == nil {
		return nil
	}

	ErrValid, ok := err.(validator.ValidationErrors)
	if !ok {
		logger.Info(err)
		return err
	}

	for _, vErr := range ErrValid {
		logger.Fatal("%s has value of %v which doesn't satisfy %s \n",
			vErr.Field(), vErr.Value(), vErr.Tag())
	}

	return nil
}

func IsName(name string) bool {
	err := Validator.Var(name, "required,min=8,containsany=!@#?")
	if err != nil {
		logger.Fatal(err, ":Invalid entered name")
		return false
	}
	return true
}

func IsEmail(email string) bool {
	err := Validator.Var(email, "required,email")
	if err != nil {
		logger.Fatal(err, " :Invalid entered email")
		return false
	}
	return true
}

func IsPassword(password string) bool {
	err := Validator.Var(password, "required,min=8,containsany=!@#?")
	if err != nil {
		logger.Fatal(err, ":Invalid entered password")
		return false
	}
	return true
}
