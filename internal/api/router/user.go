package router

import  "social_network/internal/api/v1"


func init() {
	APIRouter.HandleFunc("/home", v1.Authentication(v1.UserIndex)).Methods("GET", "POST")

	APIRouter.HandleFunc("/home/{change_name}/name", v1.Authentication(v1.UserIndex)).Methods("GET", "POST")

	APIRouter.HandleFunc("/home/{change_email}/email", v1.Authentication(v1.UserIndex)).Methods("GET", "POST")
	
	APIRouter.HandleFunc("/home/{change_password}/password", v1.Authentication(v1.UserIndex)).Methods("GET", "POST")
}
