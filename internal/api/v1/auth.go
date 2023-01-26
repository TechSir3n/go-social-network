package v1

import (
	"context"
	"fmt"
	"log"
	"net/http"
	entity "social_network/internal/domain/entities"
	"social_network/internal/repository"
	"social_network/utils"
)

func Login(wrt http.ResponseWriter, req *http.Request) {
	wrt.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func SignUp(wrt http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {

		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(wrt, "ParseForm()err:%v", err)
			return
		}

		name := req.FormValue("name")
		email := req.FormValue("email")
		password := req.FormValue("password")
		confirm_pswd := req.FormValue("confirm_pswd") // password to confirm
		fmt.Println(name, email, password, confirm_pswd)

		user := entity.User{}
		user.Name = name
		user.Password = password
		user.Email = email
		user.ConfirmPassword = confirm_pswd

		if password != confirm_pswd {
			wrt.WriteHeader(http.StatusNotFound)
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", "Password do not match")
			return
		}

		ctx := context.Background()
		id_user, err := database.CreateUser(ctx, user)

		if err != nil {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", "Failed to Create User")
			return
		}

		log.Printf("User success created: %s", id_user.ID)

		http.Redirect(wrt, req, "/login", http.StatusSeeOther)
	}

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", nil)
}

//func Auth(next http.HandlerFunc) http.HandlerFunc {}
