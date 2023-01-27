package v1

import (
	"context"
	"log"
	"net/http"
	"os"

	entity "social_network/internal/domain/entities"
	DB "social_network/internal/repository"
	"social_network/utils"

	"github.com/pkg/errors"
)


func Login(wrt http.ResponseWriter, req *http.Request) {
	wrt.Header().Set("Content-Type", "text/html; charset=utf-8")

}

func Logout(wrt http.ResponseWriter, req *http.Request) {
	wrt.Header().Set("Content-Type", "text/html; charset=utf-8")

}

func SignUp(wrt http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {

		name := req.FormValue("name")
		email := req.FormValue("email")
		password := req.FormValue("password")
		confirm_pswd := req.FormValue("confirm_pswd") // password to confirm

		if !utils.IsName(name) {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", "Wrong name entered,or user with so name already exists")
			os.Exit(1)
		} else if !utils.IsEmail(email) {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", "Wrong email entered,or so email already exists")
			os.Exit(1)
		} else if !utils.IsPassword(password) {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", "Wrong password entered")
			os.Exit(1)
		}

		user := entity.User{}
		hash, err := utils.HashPassword(user.Password)

		if err != nil {
			errors.Wrap(err, ": Failed hash password")
		}

		user.Name = name
		user.Password = string(hash)
		user.Email = email
		user.ConfirmPassword = confirm_pswd

		if password == "" && password != confirm_pswd {
			wrt.WriteHeader(http.StatusNotFound)
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", "Password do not match")
			os.Exit(1)
		}

		ctx := context.Background()
		id_user, err := DB.CreateUser(ctx, user) 

		if err != nil {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", err)
			os.Exit(1)
		}

		log.Printf("User success created: %s", id_user.ID)

		http.Redirect(wrt, req, "/login", http.StatusSeeOther)
	}

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", nil)

}

