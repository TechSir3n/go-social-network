package router

import (
	"social_network/internal/api/router"
	"social_network/internal/oauth/github"
)

func init(){
	router.APIRouter.HandleFunc("/login/github",github.GithubLogin)

	router.APIRouter.HandleFunc("/login/github/callback",github. GithubCallback)
}