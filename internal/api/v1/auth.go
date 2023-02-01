package v1

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"social_network/internal/api/v1/models"
	session "social_network/internal/domain/v2"
	"social_network/internal/pkg/jwt"
	database "social_network/internal/repository"
	"social_network/utils"
)

func Login(wrt http.ResponseWriter, req *http.Request) {
	wrt.Header().Set("Content-Type", "text/html; charset=utf-8")

	if req.Method == http.MethodPost {
		username := req.FormValue("name")
		email := req.FormValue("email")
		password := req.FormValue("password")

		apiUser := models.SignInRequest{}
		apiUser.Username = username
		apiUser.Email = email
		apiUser.Password = password

		ctx := context.Background()
		user, err := database.GetUserByEmail(ctx, email)
		if err != nil {
			wrt.WriteHeader(http.StatusNotFound)
			log.Println(err, " :User not found")
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/login.html", "User not found")
			return
		}

		compare := utils.CheckPasswordHash(user.Password, apiUser.Password)
		if !compare {
			wrt.WriteHeader(http.StatusForbidden)
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/login.html", "password is wrong")
			fmt.Println("Here error")
			return
		}

		payload, err := jwt.GenerateJWT(user) // generate access and refresh token
		if err != nil {
			log.Println(err, " :[ERROR] Generate JWT")
			return
		}

		fmt.Println("Access Token: ", payload.RefreshToken)
		fmt.Println("Access Token: ", payload.RefreshUID)
		fmt.Println("Access Token: ", payload.ExpiresRefresh)

		err = database.CreateSessions(ctx, payload) // add refresh token to database session
		if err != nil {
			log.Println(err, ": [ERROR] Create Session")
			return
		}

		//session.SetCookie(wrt, access_token) // set cookie http.Cookie(...)

		if err != nil {
			log.Println(err, ": Invalid Token")
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/login.html", "Invalid Generate token")
			return
		}

		http.Redirect(wrt, req, "/user", http.StatusSeeOther)
	}

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/login.html", nil)

}

func Authentication(endPoint http.HandlerFunc) http.HandlerFunc {
	return func(wrt http.ResponseWriter, req *http.Request) {
		wrt.Header().Set("Content-Type", "text/html; charset=utf-8")

		cookie, err := req.Cookie("Access_Token")
		access_token := cookie.Value

		if err != nil {
			log.Println(err, "Not found name cookie")
			return
		}

		value, err := jwt.ParseJWT(access_token)
		fmt.Println("Value: ", value)

		if err != nil {
			http.Error(wrt, "Token isn't valid", http.StatusUnauthorized)
			return
		}

		endPoint(wrt, req)
	}
}

func SignUp(wrt http.ResponseWriter, req *http.Request) {
	wrt.Header().Set("Content-Type", "text/html; charset=utf-8")

	if req.Method == http.MethodPost {
		name := req.FormValue("name")
		email := req.FormValue("email")
		password := req.FormValue("password")
		confirm_pswd := req.FormValue("confirm_pswd") // password to confirm

		fmt.Println("Form Password: ", password)

		if !utils.IsName(name) {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", "Wrong name entered,or user with so name already exists")
			return
		} else if !utils.IsEmail(email) {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", "Wrong email entered,or so email already exists")
			return
		} else if !utils.IsPassword(password) {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", "Wrong password entered")
			return
		}

		hash, err := utils.HashPassword(password)

		if err != nil {
			log.Println(err, ": Failed to hashing password")
			return
		}

		user := models.User{}
		user.Name = name
		user.Password = hash
		user.Email = email
		user.ConfirmPassword = confirm_pswd

		if password == "" || password != user.ConfirmPassword {
			wrt.WriteHeader(http.StatusNotFound)
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", "Password do not match")
			return
		}

		ctx := context.Background()
		id_user, err := database.CreateUser(ctx, user)

		if err != nil {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", err)
			return
		}

		log.Printf("User success created: %s", id_user.ID)

		http.Redirect(wrt, req, "/login", http.StatusSeeOther)
	}
	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/signup.html", nil)
}

func Logout(wrt http.ResponseWriter, req *http.Request) {
	session.ClearCookie(wrt)
	http.Redirect(wrt, req, "/", http.StatusSeeOther)
}

func RestorePassword(wrt http.ResponseWriter, req *http.Request) {
	wrt.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method == http.MethodPost {
		email := req.FormValue("email")
		password := req.FormValue("password")

		if email == "" && password == "" {
			http.Error(wrt, "Data claims is empty", http.StatusBadRequest)
			log.Println("Data from form is empty")
			return
		}

		ctx := context.Background()
		user, err := database.GetUserByEmail(ctx, email)

		if err != nil {
			log.Println(err, "Failed to found user with so email address")
			return
		}

		hash, err := utils.HashPassword(password)

		if err != nil {
			log.Println(err, ":Failed to hashing password")
			return
		} else if user.Password == password {
			http.Error(wrt, "Prevent changing the password to the old password", http.StatusBadRequest)
			return
		}

		user.Password = hash
		user.Email = email
		updated_data, err := database.UpdateUser(ctx, user)

		if err != nil {
			log.Println(err, "Failed to update password's user")
			return
		}

		fmt.Println(updated_data.Email, " :Success updated !")
		http.Redirect(wrt, req, "/login", http.StatusSeeOther)
	}

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/restorepassword.html", nil)
}

func ResetPassword(wrt http.ResponseWriter, req *http.Request) {

}
