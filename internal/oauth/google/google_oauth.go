package google

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	model "social_network/internal/oauth/google/model"
	"social_network/internal/repository/database/postgresql/oauth"
	"social_network/utils"
	"social_network/utils/logger"
)

var (
	OauthConfig *oauth2.Config
)

func init() {
	OauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/google/callback",
		ClientID:     "407317861817-nu5ugsv6g338rl8thjmui5or72ssga29.apps.googleusercontent.com", // os.Getenv("GOOGLE_CLIENT_ID")
		ClientSecret: "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",                                      // os.Getenv("GOOGLE_CLIENT_SECRET")
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

var rand_state = utils.GenerateRandomString()

func LoginGoogle(wrt http.ResponseWriter, req *http.Request) {
	url := OauthConfig.AuthCodeURL(rand_state)
	http.Redirect(wrt, req, url, http.StatusMovedPermanently)
}

func CallBackGoogle(wrt http.ResponseWriter, req *http.Request) {
	content, err := GetUserInfoGoogle(req.FormValue("state"), req.FormValue("code"))
	if err != nil {
		logger.Error(err.Error(), "Failed to get user data ")
	}

	var user model.GoogleContentUser
	err = json.Unmarshal(content, &user)
	if err != nil {
		logger.Error(err.Error(), "Failed to unmashall user's data turn structure")
	}

	var gl database.Google
	_, err = gl.CreateGoogleUser(context.Background(), user)
	if err != nil {
		logger.Fatal(err.Error(), "Failed to create google's user")
	}

	// after sucesss authorhization,redirect to home page
	http.Redirect(wrt, req, "/home", http.StatusSeeOther)
}

func GetUserInfoGoogle(state string, code string) ([]byte, error) {
	if state != rand_state {
		logger.Error("state is not valid")
	}

	token, err := OauthConfig.Exchange(oauth2.NoContext, code) // convert code in token
	if err != nil {
		logger.Error(err.Error(), "Failed to convert code turn token")
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken) // get response after handle token
	if err != nil {
		logger.Error(err.Error(), "Invalid get response")
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error(err.Error(), "Failed to read response.Body")
	}

	return contents, nil
}
