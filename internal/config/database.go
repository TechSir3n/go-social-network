package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
		"os"
	"github.com/joho/godotenv"
)


func Init() {
	if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Nof Found .env file")
	}
}

func ConnectDB() *pgx.Conn {

	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, username, password, dbname)

	conn, err := pgx.Connect(context.Background(),psqlInfo)
	if err != nil {
		log.Fatal("Failed open database: ", err)
	}

	defer conn.Close(context.Background())

	fmt.Println("Database connected")

	return conn
}
