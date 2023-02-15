package router

import (
	"social_network/internal/api/router/options"
	"social_network/internal/api/v1"
)

func init() {
	router.APIRouter.HandleFunc("/user/profile", v1.Authentication(v1.Profile))

	router.APIRouter.HandleFunc("/user/music", v1.Authentication(v1.Music))

	router.APIRouter.HandleFunc("/user/settings", v1.Authentication(v1.Settings))

	router.APIRouter.HandleFunc("/user/video", v1.Authentication(v1.Video))

	router.APIRouter.HandleFunc("/user/bookmarks", v1.Authentication(v1.Bookmarks))

	router.APIRouter.HandleFunc("/user/message", v1.Authentication(v1.Message))
}
