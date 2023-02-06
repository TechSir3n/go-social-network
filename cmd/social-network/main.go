package main

import (
	"log"
	"net/http"
	"os"
	"social_network/internal/api/router"
	"social_network/utils/password"
)


func main() {
	generator.GeneratePassword()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port,router.APIRouter))
}
