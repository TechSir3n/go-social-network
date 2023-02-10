package main

import (
	"net/http"
	"os"
	"social_network/internal/api/router"
   _ "social_network/internal/api/router/github"
   _ "social_network/internal/api/router/google"
	"social_network/utils/password"
	"social_network/utils/logger"
)


func main() {
	generator.GeneratePassword()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	logger.Fatal(http.ListenAndServe(":"+port,router.APIRouter))
}
