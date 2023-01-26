package main

import (
	_ "github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"social_network/internal/api/v2"
	"social_network/internal/config"
)

func main() {
	config.ConnectDB()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := v2.InitRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
