package config

import (
	"context"
	"fmt"
	"os"
	"social_network/utils/logger"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("C:/Users/Ruslan/Desktop/go-social-network/.env"); err != nil {
		logger.Error("Failed to load .env file")
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

	conn, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		logger.Error(err, ": Unable to connect database")
	}

	logger.Info("Database connected successfully")

	return conn
}
