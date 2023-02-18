package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock := NewMock()

	defer db.Close()

	mock.ExpectExec("INSERT INTO users").
		WithArgs("john@mail.ru","Join12345","Join","Join12345").
		WillReturnError(fmt.Errorf("Some error"))

	_, err:= db.Exec("INSERT INTO users(email,password,name,confirm_password) VALUES (?, ? ,? ,?)", "john@mail.ru","Join12345","Join","Join12345")

	if err != nil {
		t.Errorf("error '%s' was not expected, while inserting a row", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestGetUsers(t *testing.T) {
	_, mock := NewMock()
	mock.ExpectQuery("SELECT id,email,name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "name"}).
			AddRow("2", "email@example.ru", "Ruslan"))

	var s Postgres
	resp, err := s.GetUsers(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 5, len(resp))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not excepted when opening a stub database connection", err)
	}
	return db, mock
}
