package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"social_network/internal/api/router"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	staticDir := "/static/"
	router.APIRouter.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))


	log.Fatal(http.ListenAndServe(":"+port, router.APIRouter))
}
