package models

import "github.com/jackc/pgx/v5"

type Admin struct {
	DB          *pgx.Conn
	Special_Key string
}
