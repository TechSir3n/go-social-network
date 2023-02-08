package database

import (
	"context"
	"log"
	"social_network/internal/api/v1/models"
	"social_network/internal/config/database"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// CRUD operations Postgresql

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	ChangeEmailAddress(ctx context.Context) error

	DeleteUser(ctx context.Context, id string) (models.User, error)

	UpdateUserEmail(ctx context.Context, address string) error
	UpdateUserName(ctx context.Context) error
	UpdateUserPassword(ctx context.Context) error

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

func CreateUser(ctx context.Context, user models.User) (models.User, error) {
	user.DB = config.ConnectDB()

	user.CreatedAt = time.Now().Format(time.ANSIC)
	user.UpdatedAt = time.Now().Format(time.ANSIC)

	sqlInsert := `INSERT INTO users (email,password,name,confirm_password,created_at,updated_at) 
    VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := user.DB.Exec(ctx, sqlInsert, user.Email, user.Password, user.Name, user.ConfirmPassword, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		errors.Wrap(err, "Failed to insert data into database")
		return models.User{}, err
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
		return models.User{}, err
	}

	return models.User{
		ID: id,
	}, nil
}

func UpdateUserEmail(ctx context.Context, address, id_user string) error {
	var user models.User
	user.UpdatedAt = time.Now().Format(time.ANSIC)
	user.DB = config.ConnectDB()

	sqlUpdate := `UPDATE users SET updated_at=$1,email=$2 WHERE id=$3`
	_, err := user.DB.Exec(ctx, sqlUpdate, user.UpdatedAt, address, id_user)
	if err != nil {
		errors.Wrap(err, " :[ERROR]")
		return err
	}

	return nil
}

func UpdateUserPassword(ctx context.Context, password, id_user string) error {
	var user models.User
	user.DB = config.ConnectDB()
	user.UpdatedAt = time.Now().Format(time.ANSIC)

	sqlUpdate := `UPDATE users SET updated_at=$1,password=$2 WHERE id=$3`
	_, err := user.DB.Exec(ctx, sqlUpdate, user.UpdatedAt, password, id_user)
	if err != nil {
		log.Println(err, " :[ERROR]")
		return err
	}

	return nil
}

func UpdateUserName(ctx context.Context, username, id_user string) error {
	var user models.User
	user.DB = config.ConnectDB()
	user.UpdatedAt = time.Now().Format(time.ANSIC)
	sqlUpdate := `UPDATE users SET updated_at=$1,name=$2 WHERE id=$3`
	_, err := user.DB.Exec(ctx, sqlUpdate, user.UpdatedAt, username, id_user)
	if err != nil {
		log.Println(err, " :[ERROR]")
		return err
	}

	return nil
}
