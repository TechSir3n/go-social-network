package v1

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"social_network/internal/domain/entities"
)

// CRUD operations 

type UserService interface {
	CreateUser(user *entity.User) (entity.User, error)
	DeleteUser(id string) (entity.User, error)
	UpdateUser(user *entity.User) (entity.User, error)

	GetUser(user *entity.User) (entity.User, error)
	GetUserByID(id string) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
}

func CreateUser(user *entity.User) (entity.User, error) {

	n, err := user.DB.Exec(context.Background(), "INSERT INTO users(email,password),VALUES($1,$2)", user.Email, user.Password)

	if err != nil {
		errors.Wrap(err, "Faied to insert data user to database")
	}

	if n.RowsAffected() == 0 {
		fmt.Errorf("Nothing has been inserted into the database")
	}

	return entity.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func DeleteUser(id string) (entity.User, error) {
	var user entity.User
	n, err := user.DB.Exec(context.Background(), "DELETE FROM users WHERE id =$1", id)

	if err != nil {
		errors.Wrap(err, "Failed to delete user,incorrect id or user with such id doesn't exists")
	}

	if n.RowsAffected() == 0 {
		fmt.Errorf("Nothing has been removed from the database")
	}

	return entity.User{
		ID: user.ID,
	}, nil
}

func UpdateUser(user *entity.User) (entity.User, error) {
	// need to end a bit later 
	n, err := user.DB.Exec(context.Background(), "UPDATE users SET email=$1,password=$2 WHERE email")

	if err != nil {
		errors.Wrap(err, "Failed to update user,incorrect enter data")
	}

	if n.RowsAffected() == 0 {
		fmt.Errorf("Nothing has been deleted into the database")
	}

	return entity.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func GetUser(user *entity.User) (*entity.User, error) {
	rows, err := user.DB.Query(context.Background(), "SELECT id_user, email,password FROM users")

	if err != nil {
		errors.Wrap(err, "Failed to get some data of the database")
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(user.ID, user.Email, user.Password)
	}

	if rows.Err() != nil {
		errors.Wrap(rows.Err(), "Failed to read data of the database")
	}

	return &entity.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func GetUserByID(id string) (entity.User, error) {

}

func GetUserByEmail(email string) (entity.User, error) {

}
