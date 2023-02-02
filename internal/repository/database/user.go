package database 

import (
	"social_network/internal/api/v1/models"
	"github.com/pkg/errors"
	"social_network/internal/config/database"
	"context"
	"log"
)

func GetUser(ctx context.Context) ([]models.User, error) {
	var user models.User
	user.DB = config.ConnectDB()
	rows, err := user.DB.Query(ctx, "SELECT email,password,name,confirm_password FROM users")
	if err != nil {
		errors.Wrap(err, "Failed to get some data from the database")
		return []models.User{}, err
	}

	defer rows.Close()

	data := []models.User{}

	for rows.Next() {
		err := rows.Scan(user.ID, user.Email, user.Password)
		if err != nil {
			errors.Wrap(err, " :[ERROR]")
		}

		data = append(data, user)
	}

	if rows.Err() != nil {
		errors.Wrap(rows.Err(), " :[ERROR]")
		return []models.User{}, err
	}

	return data, nil
}


func GetUserByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	user.DB = config.ConnectDB()
	rows, err := user.DB.Query(ctx, "SELECT * FROM users WHERE id_user=$1", id)

	if err != nil {
		errors.Wrap(err, "Couldn't be found data with such id into database")
		return models.User{}, err
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.Email, &user.Password)
	}

	if rows.Err() != nil {
		errors.Wrap(rows.Err(), " :[ERROR]")
		return models.User{}, err
	}

	return models.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	user.DB = config.ConnectDB()
	rows, err := user.DB.Query(ctx, "SELECT email,password,name FROM users WHERE email = $1", email)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&user.Email, &user.Password, &user.Name); err != nil {
			log.Fatal(err)
		}
	}

	return models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
