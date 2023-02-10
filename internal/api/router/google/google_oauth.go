package router 

import (
	"social_network/internal/api/router"
	"social_network/internal/oauth/google"
)

func init(){
	router.APIRouter.HandleFunc("/google/login",google.LoginGoogle)

	router.APIRouter.HandleFunc("/google/callback",google.CallBackGoogle)
}