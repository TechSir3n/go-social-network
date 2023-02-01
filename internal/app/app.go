package app

import (
	"log"
	"net/http"
	"os"

	"social_network/internal/api/router"
	"social_network/internal/config"
)

func Run() {
	config.InitRedis()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.APIRouter.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Fatal(http.ListenAndServe(":"+port, router.APIRouter))
}
