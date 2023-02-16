package github

import (
	"bytes"
	 "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"social_network/internal/api/v1"
	_ "social_network/internal/config/database"
	model "social_network/internal/oauth/github/model"
	oauth "social_network/internal/repository/database/postgresql/oauth"
	"social_network/utils"
	"social_network/utils/logger"
)

var rand_state = utils.GenerateRandomString()

func GithubLogin(wrt http.ResponseWriter, req *http.Request) {
	cliendID := GetGithubClientID()
	redirectURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&state=%s", cliendID, "http://localhost:3000/login/github/callback", rand_state)
	http.Redirect(wrt, req, redirectURL, http.StatusMovedPermanently)
}

func GithubCallback(wrt http.ResponseWriter, req *http.Request) {
	state := req.URL.Query().Get("state")
	code := req.URL.Query().Get("code")             // get code of url in which is access token
	githubAccessToken := GetGithubAccessToken(code) // add access token to get user's data
	githubData := GetGithubData(state,githubAccessToken)  // get user's data

	var github oauth.GitHub
	github.GitHubUser.CreateGitHubUser(context.Background(), githubData)
	v1.Home(wrt, req)
}

func GetGithubData(state,accessToken string) model.GitHubUserDataResponse {
	if state!=rand_state{
		logger.Error("Something went wrong, CSRF Attack may be")
	}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		logger.Panic("API Request creation failed")
	}

	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Panic("Request failed")
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	var data model.GitHubUserDataResponse
	err = json.Unmarshal(respbody, &data)
	if err != nil {
		logger.Fatal("Failed umarshal body into structure", err.Error())
	}
	

	return data
}

func GetGithubAccessToken(code string) string {

	clientID := GetGithubClientID()
	clientSecret := GetGithubClientSecret()

	requestBody := model.RequestBodyGitHub{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Code:         code,
	}
	requestJSON, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(requestJSON))
	if err != nil {
		logger.Panic("Request creation failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Panic("Request failed")
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	var ghresp model.GitHubAccessTokenResponse
	json.Unmarshal(respbody, &ghresp)

	return ghresp.AccessToken
}

func GetGithubClientID() string {

	clientID, exists := os.LookupEnv("GITHUB_CLIENT_ID")
	if !exists {
		logger.Fatal("Github Client ID not defined in .env file")
	}

	return clientID
}

func GetGithubClientSecret() string {

	clientSecret, exists := os.LookupEnv("GITHUB_CLIENT_SECRET")
	if !exists {
		logger.Fatal("Github Client ID not defined in .env file")
	}

	return clientSecret
}
