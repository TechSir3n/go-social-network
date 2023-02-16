package v1

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"social_network/utils"
)

func UserIndex(wrt http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	upd_pass := vars["change_password"]
	upd_name := vars["change_name"]
	upd_email := vars["change_email"]

	update_user := req.FormValue("")

	if upd_name == "update_name" {
		UpdateUserName(wrt, req, update_user)
	} else if upd_pass == "update_password" {
		UpdateUserPassword(wrt, req, update_user)
	} else if upd_email == "update_email" {
		UpdateUserEmail(wrt, req, "")
	}

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/home/html/settings.html", nil)
}

func UpdateUserName(wrt http.ResponseWriter, req *http.Request, name string) {
	if req.Method == http.MethodPost {
		id_user := req.FormValue("search")
		ctx := context.Background()
		err := db.User.UpdateUserName(ctx, name, id_user)
		if err != nil {
			errors.Wrap(err, " :FAILED to update username")
			return
		}
	}
}

func UpdateUserPassword(wrt http.ResponseWriter, req *http.Request, password string) {
	if req.Method == http.MethodPost {
		id_user := req.FormValue("search")
		ctx := context.Background()
		err := db.User.UpdateUserPassword(ctx, password, id_user)
		if err != nil {
			errors.Wrap(err, " :FAILED to update password")
			return
		}
	}

}

func UpdateUserEmail(wrt http.ResponseWriter, req *http.Request, email string) {
	if req.Method == http.MethodPost {
		id_user := req.FormValue("")
		ctx := context.Background()
		err := db.User.UpdateUserEmail(ctx, email, id_user)
		if err != nil {
			errors.Wrap(err, " :FAILED to update email")
			return
		}
	}

}
