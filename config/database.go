package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *pgx.Conn {

   func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Nof Found .env file,err: ", err)
		}
	}()

	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)

	conn, err := pgx.Connect(context.Background(), psqlInfo)

	if err != nil {
		log.Fatal("Connect err: ", err)
	}

	fmt.Println("Database connected")

	defer conn.Close(context.Background())

	return conn
}
