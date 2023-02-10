package database

import (
	"context"
	"social_network/internal/config/database"
	"social_network/internal/oauth/google/model"
	"social_network/utils/logger"
)

type GoogleRepository interface {
	CreateGoogleUser(ctx context.Context, user google.GoogleContentUser) ([]google.GoogleContentUser, error)
	GetGoogleUsers(ctx context.Context) ([]GoogleRepository, error)
	GetGoogleUserByEmail(ctx context.Context) (google.GoogleContentUser, error)

	UpdateGoogleUser(ctx context.Context, user google.GoogleContentUser, name string) error
	DeleteGoogleUser(ctx context.Context, name string) error
}

type GoogleUser struct {
	GoogleRepository GoogleRepository
	GoogleUser       google.GoogleContentUser
}

func NewGoogleUser(GoogleRep GoogleRepository, model google.GoogleContentUser) *GoogleUser {
	return &GoogleUser{
		GoogleRepository: GoogleRep,
		GoogleUser:       model,
	}
}

func CreateGoogleUser(ctx context.Context, user google.GoogleContentUser) (google.GoogleContentUser, error) {
	db := config.ConnectDB()

	sqlInsert := `INSERT INTO GoogleContentUser (id,email) 
    VALUES ($1,$2)`
	_, err := db.Exec(ctx, sqlInsert, user.ID, user.Email)

	if err != nil {
		logger.Panic("Unable insert into database [GoogleContentUser]", err.Error())
		return google.GoogleContentUser{}, err
	}

	return google.GoogleContentUser{}, err
}

func GetGoogleUsers(ctx context.Context) ([]google.GoogleContentUser, error) {
	var user google.GoogleContentUser
	db := config.ConnectDB()

	rows, err := db.Query(ctx, "SELECT id,email FROM GoogleContentUser")
	if err != nil {
		logger.Error("Unable get data from database [GoogleContentUser]", err.Error())
	}

	defer rows.Close()

	data := []google.GoogleContentUser{}

	for rows.Next() {
		if err = rows.Scan(); err != nil {
			logger.Error("Failed to scan into structure [GoogleContentUser]", err.Error())
		}
		data = append(data, user)
	}

	return data, nil
}

func GetGoogleUserByEmail(ctx context.Context, email string) (google.GoogleContentUser, error) {
	var user google.GoogleContentUser
	db := config.ConnectDB()

	row := db.QueryRow(ctx, "SELECT id,email FROM  GitHubUserData WHERE email=$1", email)

	if err := row.Scan(); err != nil {
		logger.Error(err.Error(), " : Failed to scan into structure[GoogleContentUser]")
		return google.GoogleContentUser{}, err
	}

	return user, nil
}

func UpdateGoogleUser(ctx context.Context, user google.GoogleContentUser, name string) error {
	db := config.ConnectDB()

	sqlInsert := `UPDATE GoogleContentUser SET id=$1 WHERE email=$2
    VALUES ($1,$2)`
	_, err := db.Exec(ctx, sqlInsert, user.ID, user.Email)

	if err != nil {
		logger.Panic("Unable update database [GoogleContentUser]", err.Error())
		return err
	}

	return nil
}

func DeleteGoogleUser(ctx context.Context, email string) error {
	db := config.ConnectDB()

	sqlInsert := `DELETE FROM GoogleContentUser WHERE email=$1`
	_, err := db.Exec(ctx, sqlInsert, email)
	if err != nil {
		logger.Panic("Unable Delete from database [GoogleContentUser]", err.Error())
		return err
	}

	return nil
}
