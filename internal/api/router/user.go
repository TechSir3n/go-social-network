package router

import v1 "social_network/internal/api/v1"


func init() {
	APIRouter.HandleFunc("/user", v1.Authentication(v1.UserIndex)).Methods("GET", "POST")
	
	APIRouter.HandleFunc("/user", v1.Authentication(v1.AdminIndex)).Methods("GET", "POST")
}
