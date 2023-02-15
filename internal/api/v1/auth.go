package v1

import (
	"context"
	"log"
	"net/http"
	"github.com/pkg/errors"

	"social_network/internal/api/v1/models"
	session "social_network/internal/domain/v2"
	"social_network/internal/pkg/crypt"
	"social_network/internal/pkg/jwt"
	"social_network/internal/repository/database/postgresql"
	redis "social_network/internal/repository/database/redis"
	"social_network/internal/repository/memcached"
	"social_network/utils"
	"social_network/utils/logger"
	valid "social_network/utils/validator"
)

func Login(wrt http.ResponseWriter, req *http.Request) {
	wrt.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method == http.MethodPost {
		username := req.FormValue("name")
		email := req.FormValue("email")
		password := req.FormValue("password")

		apiUser := models.SignInRequest{
			Username: username,
			Email:    email,
			Password: password,
		}

		ctx := context.Background()
		var user models.User
		user_cached, err := memcached.GetMemcached("credentials")
		//  if the user is still in the cache, then he will be cured from there otherwise from the database
		if err != nil {
			user_db, err := database.GetUserByEmail(ctx, email)
			if err != nil {
				wrt.WriteHeader(http.StatusNotFound)
				log.Println(err, " :User not found")
				utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/html/login.html", "User not found")
				return
			}
			user = user_db
		} else {
			user = user_cached
		}

		compare := crypt.CheckPasswordHash(user.Password, apiUser.Password)
		if !compare {
			wrt.WriteHeader(http.StatusForbidden)
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/html/login.html", "password is wrong")
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

		http.Redirect(wrt, req, "/home", http.StatusSeeOther)
	}

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/html/login.html", nil)

}

func Authentication(EndPoint http.HandlerFunc) http.HandlerFunc {
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

		EndPoint(wrt, req)
	}
}

func SignUp(wrt http.ResponseWriter, req *http.Request) {
	wrt.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method == http.MethodPost {
		name := req.FormValue("name")
		email := req.FormValue("email")
		password := req.FormValue("password")
		confirm_pswd := req.FormValue("confirm_pswd") // password to confirm

		if !valid.IsPassword(password) && !valid.IsEmail(email) && !valid.IsName(name) {
			http.Error(wrt, ": Invalid entered credentials", http.StatusBadRequest)
			log.Println(" :user entered invalid data")
			return
		}

		hash, err := crypt.HashPassword(password)
		if err != nil {
			log.Println(err, ": Failed to hashing password")
			return
		}

		user := models.User{
			Name:            name,
			Password:        hash,
			Email:           email,
			ConfirmPassword: confirm_pswd,
		}

		if err := memcached.SetMemcached(user); err != nil {
			logger.Info(err.Error())
		}

		if password == "" || password != user.ConfirmPassword {
			wrt.WriteHeader(http.StatusNotFound)
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/html/signup.html", "Password do not match")
			return
		}

		ctx := context.Background()
		id_user, err := database.CreateUser(ctx, user)

		if err != nil {
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/html/signup.html", err)
			return
		}

		log.Printf("User success created: %s", id_user.ID)

		http.Redirect(wrt, req, "/login", http.StatusSeeOther)
	}
	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/html/signup.html", nil)
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

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/html/restore_password.html", nil)
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
		err = database.UpdateUserPassword(ctx, user.Password, "2")
		if err != nil {
			log.Println(err, " :Failed to update user password")
			return
		}
		http.Redirect(wrt, req, "/login", http.StatusSeeOther)
	}
	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/html/reset_password.html", nil)
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
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/hmtl/login_admin.html", http.StatusForbidden)
			log.Println("Access is denied, Wrong Admin Key")
			return
		}

		http.Redirect(wrt, req, "/admin", http.StatusSeeOther)
	}

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/access/html/login_admin.html", nil)
}
