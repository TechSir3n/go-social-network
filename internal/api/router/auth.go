package router

import (
	"github.com/gorilla/mux"
	"social_network/internal/api/v1"
)

var APIRouter = mux.NewRouter()

func init() {
	

	APIRouter.HandleFunc("/registration", v1.SignUp).Methods("GET", "POST")

	APIRouter.HandleFunc("/login", v1.Login).Methods("GET", "POST")

	APIRouter.HandleFunc("/logout",v1.Logout).Methods("GET", "POST")
	
}
