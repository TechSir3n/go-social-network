package database

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"social_network/internal/api/v1/models"
)

func (s Postgres) GetUsers(ctx context.Context) ([]models.User, error) {
	var user models.User
	rows, err := user.DB.Query(ctx, "SELECT id,email,name FROM users")
	if err != nil {
		errors.Wrap(err, "Failed to get some data from the database")
		return []models.User{}, err
	}

	defer rows.Close()

	data := []models.User{}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Name)
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

func (s Postgres) GetUserByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	rows, err := user.DB.Query(ctx, "SELECT id,email,name FROM users WHERE id=$1", id)

	if err != nil {
		errors.Wrap(err, "Couldn't be found data with such id into database")
		return models.User{}, err
	}

	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Email, &user.Name); err != nil {
			log.Println(err, ": FAILED to scan into structure")
		}
	}

	if rows.Err() != nil {
		errors.Wrap(rows.Err(), " :[ERROR]")
		return models.User{}, err
	}

	return models.User{
		ID:       user.ID,
		Password: user.Name,
		Email:    user.Email,
	}, nil
}

func (s Postgres) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	rows, err := user.DB.Query(ctx, "SELECT id,email,password,name FROM users WHERE email = $1", email)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Email, &user.Name); err != nil {
			log.Println(err, ": FAILED to scan into structure")
		}
	}

	return models.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
	}, nil
}
