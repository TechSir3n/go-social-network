package v1

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"

	"social_network/utils"
)

// this method can use only admin,who have special_key so it update every 10 minutes in redis
func AdminIndex(wrt http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	delete_user := vars["method_delete"]
	get_user := vars["method_get"]
	get_all_user := vars["method_get_all"]
	find_user := vars["method_find"]

	if delete_user == "delete" {
		DeleteUser(wrt, req)
	} else if find_user == "find_user" {
		FindUser(wrt, req)
	} else if get_user == "by_id_get" {
		GetUser(wrt, req, get_user)
	} else if get_user == "by_email_get" {
		GetUser(wrt, req, get_user)
	} else if get_all_user == "get_all_users" {
		GetAllUsers(wrt, req)
	} else {
		http.Error(wrt, "Use invalid variable", http.StatusExpectationFailed)
		return
	}

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/admin/html/admin_page.html", nil)
}

func GetUser(wrt http.ResponseWriter, req *http.Request, txt string) {
	if req.Method == http.MethodPost {
		by_user := req.FormValue("search")
		ctx := context.Background()
		if strings.HasPrefix(txt, "by_id_get") {
			user_data, err := db.User.GetUserByID(ctx, by_user)
			if err != nil {
				errors.Wrap(err, " :FAILED to get user by id")
				return
			}
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/admin/html/show_users.html", user_data.Email)
		} else if strings.HasPrefix(txt, "by_email_get") {
			user_data, err := db.User.GetUserByEmail(ctx, by_user)
			if err != nil {
				errors.Wrap(err, " :FAILED to get user by id")
				return
			}
			utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/admin/html/show_users.html", user_data)
		} else {
			log.Println("none of the conditions worked")
			return
		}
	}
}

func DeleteUser(wrt http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodDelete {
		id_user := req.FormValue("search")
		ctx := context.Background()
		_, err := db.User.DeleteUser(ctx, id_user)
		if err != nil {
			errors.Wrap(err, " :FAILED to delete user")
			return
		}
	}
}

func GetAllUsers(wrt http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		ctx := context.Background()
		users, err := db.User.GetUsers(ctx)
		if err != nil {
			errors.Wrap(err, " :FAILED to get users")
			return
		}
		utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/admin/html/show_users.html", users)
	}
}

func FindUser(wrt http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {

	}
}
