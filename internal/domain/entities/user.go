package entity 

import "github.com/jackc/pgx/v5"

type User struct {
    DB *pgx.Conn
	ID       string
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

