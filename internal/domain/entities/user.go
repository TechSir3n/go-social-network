package entity 

import (
	"github.com/jackc/pgx/v5"
)

type User struct {
    DB *pgx.Conn
	ID       string
	Name     string          	`json:"Name"`
	Email    string 		 	`json:"Email"`
	Password string 		 	`json:"Password"`
	ConfirmPassword string  	`json:"ConfirmPassword"`
	CreatedAt       string      `json:"CreatedAt"`
	UpdatedAt     	string 	    `json:"UpdatedAt"`
}

