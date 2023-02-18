package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"social_network/utils/logger"
	"testing"
)

func TestConnectDB(t *testing.T) {
	dbPool, err := pgxpool.Connect(context.Background(),
		fmt.Sprintf("postgres://%v:%v@%v:%v/%v", "postgres", "password", "localhost", 5432, "postgres"))
	if err != nil {
		t.Fatal("Error while connecting to postgrs", err)
	}

	var temp string
	err = dbPool.QueryRow(context.Background(), "select 'Hello World!'").Scan(&temp)
	if err != nil {
		t.Fatal("Error while copy to variable value", err)
	}

	if temp != "Hello World!" {
		t.Fatal("Incorrect result,got:", temp)
	}

	logger.Info("Successfuly connected to Postgres")
}
