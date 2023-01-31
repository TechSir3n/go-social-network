package database

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"log"
	"social_network/internal/api/v1/models"
	"social_network/internal/config"
	"time"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUser(ctx context.Context, id string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)

	GetUser(ctx context.Context) (models.User, error)
	GetUserByID(ctx context.Context, id string) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

type UserService struct {
	UserRepository UserRepository
	User           models.User
}

func NewUserService(user UserRepository, model models.User) *UserService {
	return &UserService{
		UserRepository: user,
		User:           model,
	}
}

// CRUD operations

func CreateUser(ctx context.Context, user models.User) (models.User, error) {
	user.DB = config.ConnectDB()

	user.CreatedAt = time.Now().Format(time.ANSIC)
	user.UpdatedAt = time.Now().Format(time.ANSIC)

	sqlInsert := `INSERT INTO users (email,password,name,confirm_password,created_at,updated_at) 
    VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := user.DB.Exec(ctx, sqlInsert, user.Email, user.Password, user.Name, user.ConfirmPassword, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		errors.Wrap(err, "Failed to insert data into database")
	}

	return models.User{
		ID: user.ID,
	}, nil
}

func DeleteUser(ctx context.Context, id string) (models.User, error) {
	var user models.User
	user.DB = config.ConnectDB()
	sqlDelete := `DELETE FROM users WHERE id =$1`
	_, err := user.DB.Exec(ctx, sqlDelete, id)

	if err != nil {
		errors.Wrap(err, "Failed to delete user,incorrect id or user with such id doesn't exists")
	}

	return models.User{
		ID: id,
	}, nil
}

func UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	// need to end a bit later
	sqlUpdate := `UPDATE users SET email=$1,password=$2 WHERE email`
	user.UpdatedAt = time.Now().Format(time.ANSIC)
	_, err := user.DB.Exec(ctx, sqlUpdate) // do

	if err != nil {
		errors.Wrap(err, "Failed to update user,incorrect enter data")
	}

	return models.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func GetUser(ctx context.Context) ([]models.User, error) {
	var user models.User
	user.DB = config.ConnectDB()
	rows, err := user.DB.Query(ctx, "SELECT email,password,name,confirm_password FROM users")
	if err != nil {
		errors.Wrap(err, "Failed to get some data from the database")
	}

	defer rows.Close()

	data := []models.User{}

	for rows.Next() {
		err := rows.Scan(user.ID, user.Email, user.Password)
		if err != nil {
			errors.Wrap(err, "Failed to copy data to structure")
		}

		data = append(data, user)
	}

	if rows.Err() != nil {
		errors.Wrap(rows.Err(), "Failed to read data from the database")
	}

	return data, nil
}

func GetUserByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	user.DB = config.ConnectDB()
	rows, err := user.DB.Query(ctx, "SELECT * FROM users WHERE id_user=$1", id)

	if err != nil {
		errors.Wrap(err, "Couldn't be found data with such id into database")
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.Email, &user.Password)
	}

	if rows.Err() != nil {
		errors.Wrap(rows.Err(), "Request Rows.Nex() gave an error")
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
