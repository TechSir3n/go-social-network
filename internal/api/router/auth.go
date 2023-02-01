package router

import (
	v1 "social_network/internal/api/v1"

	"github.com/gorilla/mux"
)

var APIRouter = mux.NewRouter()

func init() {
	APIRouter.HandleFunc("/registration", v1.SignUp).Methods("GET", "POST")

	APIRouter.HandleFunc("/login", v1.Login).Methods("GET", "POST")

	APIRouter.HandleFunc("/logout", v1.Logout).Methods("GET", "POST")

	APIRouter.HandleFunc("/restore/password", v1.RestorePassword).Methods("GET", "POST")

	APIRouter.HandleFunc("/reset/password", v1.ResetPassword).Methods("GET", "POST")
}
