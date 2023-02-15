package router

import (
	"social_network/internal/api/v1"
	"social_network/internal/api/router/options"
)


func init() {
	
	router.APIRouter.HandleFunc("/registration", v1.SignUp)

	router.APIRouter.HandleFunc("/login", v1.Login)

	router.APIRouter.HandleFunc("/logout", v1.Logout)

	router.APIRouter.HandleFunc("/verify", v1.VerifyEmail)

	router.APIRouter.HandleFunc("/reset/password", v1.ResetPassword)

	router.APIRouter.HandleFunc("/access/admin", v1.AccessAdmin)
}
