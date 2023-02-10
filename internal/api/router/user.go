package router

import  "social_network/internal/api/v1"


func init() {
	APIRouter.HandleFunc("/user", v1.Authentication(v1.UserIndex)).Methods("GET", "POST")

	APIRouter.HandleFunc("/user/{change_name}/name", v1.Authentication(v1.UserIndex)).Methods("PUT")

	APIRouter.HandleFunc("/user/{change_email}/email", v1.Authentication(v1.UserIndex)).Methods("PUT")
	
	APIRouter.HandleFunc("/user/{change_password}/password", v1.Authentication(v1.UserIndex)).Methods("PUT")
}
