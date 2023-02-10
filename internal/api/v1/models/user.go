package models

import (
	"github.com/jackc/pgx/v5"
)

type User struct {
	DB              *pgx.Conn
	ID              string `json:"id"`
	Name            string `json:"name" validate:"required,min=6"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,containsany=!@#?"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8"`
	CreatedAt       string `json:"createt_at"`
	UpdatedAt       string `json:"updated_at"`
}
