package database

import (
	"context"
	"social_network/internal/config/database"
	"social_network/internal/oauth/github/model"
	"social_network/utils/logger"
)

type GitHubRepository interface {
	CreateGitHubUser(ctx context.Context, user github.GitHubUserDataResponse) ([]github.GitHubUserDataResponse, error)
	GetGitHubUsers(ctx context.Context) ([]github.GitHubUserDataResponse, error)
	GetGitHubUserByName(ctx context.Context) (github.GitHubUserDataResponse, error)

	UpdateGitHubUser(ctx context.Context, user github.GitHubUserDataResponse, name string) error
	DeleteGitHubUser(ctx context.Context, name string) error
}

type GitHubUser struct {
	GitHubRepository GitHubRepository
	GitHubUser       github.GitHubUserDataResponse
}

func NewGitHubUser(GitHubRep GitHubRepository, model github.GitHubUserDataResponse) *GitHubUser {
	return &GitHubUser{
		GitHubRepository: GitHubRep,
		GitHubUser:       model,
	}
}

func CreateGitHubUser(ctx context.Context, user github.GitHubUserDataResponse) (github.GitHubUserDataResponse, error) {
	db := config.ConnectDB()

	sqlInsert := `INSERT INTO GitHubUserData (name,login,id,location,created_at,updated_at) 
    VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := db.Exec(ctx, sqlInsert, user.Name, user.Login, user.ID, user.Location, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		logger.Panic("Unable insert into database [GitHubUserData]", err.Error())
		return github.GitHubUserDataResponse{}, err
	}

	return github.GitHubUserDataResponse{
		ID:    user.ID,
		Name:  user.Name,
		Login: user.Login,
	}, nil
}

func GetGitHubUsers(ctx context.Context) ([]github.GitHubUserDataResponse, error) {
	var user github.GitHubUserDataResponse
	db := config.ConnectDB()

	rows, err := db.Query(ctx, "SELECT (name,login,id,location,created_at,updated_at FROM  GitHubUserData")
	if err != nil {
		logger.Error("Unable get data from database [GitHubUserData]", err.Error())
	}

	defer rows.Close()

	data := []github.GitHubUserDataResponse{}

	for rows.Next() {
		if err = rows.Scan(&user.Name, &user.Login, &user.ID, &user.Location, &user.CreatedAt, &user.UpdatedAt); err != nil {
			logger.Error("Failed to scan into structure [GitHubUserData]", err.Error())
		}
		data = append(data, user)
	}

	return data, nil
}

func GetGitHubUserByName(ctx context.Context, name string) (github.GitHubUserDataResponse, error) {
	var user github.GitHubUserDataResponse
	db := config.ConnectDB()

	row := db.QueryRow(ctx, "SELECT name,login,id,location,created_at,updated_at FROM  GitHubUserData WHERE name=$1", name)

	if err := row.Scan(&user.Name, &user.Login, &user.ID, &user.Location, &user.CreatedAt, &user.UpdatedAt); err != nil {
		logger.Error(err.Error(), " : Failed to scan into structure[GitHubUserData]")
		return github.GitHubUserDataResponse{}, err
	}

	return user, nil
}

func UpdateGitHubUser(ctx context.Context, user github.GitHubUserDataResponse, name string) error {
	db := config.ConnectDB()

	sqlInsert := `UPDATE GitHubUserData SET name=$1,login=$2,id=$3,location=$4,updated_at=$5 WHERE name=$6)
    VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := db.Exec(ctx, sqlInsert, user.Name, user.Login, user.ID, user.Location, user.UpdatedAt, name)

	if err != nil {
		logger.Panic("Unable update database [GitHubUserData]", err.Error())
		return err
	}

	return nil
}

func DeleteGitHubUser(ctx context.Context, name string) error {
	db := config.ConnectDB()

	sqlInsert := `DELETE FROM GitHubUserData WHERE name=$1`
	_, err := db.Exec(ctx, sqlInsert, name)
	if err != nil {
		logger.Panic("Unable Delete from database [GitHubUserData]", err.Error())
		return err
	}

	return nil
}
