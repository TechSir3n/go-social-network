package entity 

import "database/sql"

type User struct {
    DB *sql.DB
	ID       string
	Name     string          `json:"Name"`
	Email    string 		 `json:"Email"`
	Password string 		 `json:"Password"`
	ConfirmPassword string   `json:"ConfirmPassword"`
}

