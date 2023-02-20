package router

import (
 "social_network/internal/api/v1"
 "social_network/internal/api/router/options"
)


func init() {

	router.APIRouter.HandleFunc("/user/{change_name}/name", v1.Authentication(v1.UserIndex))

	router.APIRouter.HandleFunc("/user/{change_email}/email", v1.Authentication(v1.UserIndex))
	
	router.APIRouter.HandleFunc("/user/{change_password}/password", v1.Authentication(v1.UserIndex))
}
