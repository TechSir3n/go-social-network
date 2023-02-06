package v1

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"net/http"

	"social_network/internal/api/v1/models"
	session "social_network/internal/domain/v2"
	"social_network/internal/pkg/crypt"
	"social_network/internal/pkg/jwt"
	"social_network/internal/repository/database/postgresql"
	redis "social_network/internal/repository/database/redis"
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
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/login.html", "User not found")
			return
		}

		compare := crypt.CheckPasswordHash(user.Password, apiUser.Password)
		if !compare {
			wrt.WriteHeader(http.StatusForbidden)
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/login.html", "password is wrong")
			return
		}

		payload, err := jwt.GenerateJWT(user) // generate access and refresh token
		if err != nil {
			log.Println(err, " :[ERROR] Generate JWT")
			return
		}

		err = database.CreateSessions(ctx, payload) // add refresh token to database session
		if err != nil {
			log.Println(err, ": [ERROR] Create Session")
			return
		}

		session.SetCookie(wrt, payload.AccessToken) // set cookie http.Cookie(...)

		if err != nil {
			log.Println(err, ": Invalid Token")
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/login.html", "Invalid Generate token")
			return
		}

		http.Redirect(wrt, req, "/user", http.StatusSeeOther)
	}

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/login.html", nil)

}

func Authentication(endPoint http.HandlerFunc) http.HandlerFunc {
	return func(wrt http.ResponseWriter, req *http.Request) {
		wrt.Header().Set("Content-Type", "text/html; charset=utf-8")

		cookie, err := req.Cookie("Session")
		access_token := cookie.Value

		if err != nil {
			log.Println(err, "Not found name cookie")
			return
		}

		_, err = jwt.ParseJWT(access_token)

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

		if !utils.IsName(name) {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/signup.html", "Wrong name entered,or user with so name already exists")
			return
		} else if !utils.IsEmail(email) {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/signup.html", "Wrong email entered,or so email already exists")
			return
		} else if !utils.IsPassword(password) {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/signup.html", "Wrong password entered")
			return
		}

		hash, err := crypt.HashPassword(password)

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
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/signup.html", "Password do not match")
			return
		}

		ctx := context.Background()
		id_user, err := database.CreateUser(ctx, user)

		if err != nil {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/signup.html", err)
			return
		}

		log.Printf("User success created: %s", id_user.ID)

		http.Redirect(wrt, req, "/login", http.StatusSeeOther)
	}
	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/signup.html", nil)
}

func Logout(wrt http.ResponseWriter, req *http.Request) {
	session.ClearCookie(wrt)
	http.Redirect(wrt, req, "/", http.StatusSeeOther)
}

func VerifyEmail(wrt http.ResponseWriter, req *http.Request) {
	wrt.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method == http.MethodPost {
		email := req.FormValue("email")

		if email == "" {
			http.Error(wrt, "Data claims is empty", http.StatusBadRequest)
			log.Println("Data from form is empty")
			return
		}

		ctx := context.Background()
		_, err := database.GetUserByEmail(ctx, email)
		if err != nil {
			http.Error(wrt, "User not found", http.StatusBadRequest)
			log.Println(err, "Failed to found user with so email address")
			return
		}

		http.Redirect(wrt, req, "/reset/password", http.StatusSeeOther)
	}

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/restore_password.html", nil)
}

func ResetPassword(wrt http.ResponseWriter, req *http.Request) {
	wrt.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method == http.MethodPost {
		password := req.FormValue("password")
		if password == "" {
			http.Error(wrt, "Password must be filled ", http.StatusBadRequest)
			return
		}

		var user models.User
		hash, err := crypt.HashPassword(password)
		user.Password = hash

		if err != nil {
			log.Println(err, " :Failed to hashing password")
			return
		}

		ctx := context.Background()
		err = database.UpdateUserPassword(ctx, user.Password)
		if err != nil {
			log.Println(err, " :Failed to update user password")
			return
		}
		http.Redirect(wrt, req, "/login", http.StatusSeeOther)
	}
	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/reset_password.html", nil)
}

// login for administrator,using a special key
func AccessAdmin(wrt http.ResponseWriter, req *http.Request) {
	wrt.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method == http.MethodPost {
		key := req.FormValue("special_key")
		if key == "" {
			http.Error(wrt, "Key must be written", http.StatusBadRequest)
			return
		}

		ctx := context.Background()
		admin, err := redis.GetAdminPassword(ctx)
		if err != nil {
			errors.Wrap(err, " Unable get :[ADMIN]")
			return
		}
	
		if admin.Special_Key != key {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/admin.html", http.StatusForbidden)
			log.Println("Access is denied,Wrong Admin Key")
			return
		}
	}
	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/admin.html", nil)
}
