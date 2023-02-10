package router 

import (
   "social_network/internal/api/v1"
)

func init() {
	APIRouter.HandleFunc("/home", v1.Authentication(v1.Home)).Methods("GET", "POST")

	APIRouter.HandleFunc("/home/settings", v1.Authentication(v1.Settings)).Methods("GET", "POST")

	APIRouter.HandleFunc("/home/music", v1.Authentication(v1.Music)).Methods("GET", "POST")

	APIRouter.HandleFunc("/home/video", v1.Authentication(v1.Video)).Methods("GET", "POST")

	APIRouter.HandleFunc("/home/bookmarks", v1.Authentication(v1.Bookmarks)).Methods("GET", "POST")

	APIRouter.HandleFunc("/home/message", v1.Authentication(v1.Message)).Methods("GET", "POST")
}


