package models


type SignInRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"Role"`
}

type SignInResponse struct {
	AccessToken string `json:"access_toekn"`
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}


