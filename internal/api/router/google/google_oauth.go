package router 

import (
	"social_network/internal/api/router/options"
	"social_network/internal/oauth/google"
)

func init(){
	router.APIRouter.HandleFunc("/login/google",google.LoginGoogle)

	router.APIRouter.HandleFunc("/google/callback",google.CallBackGoogle)
}