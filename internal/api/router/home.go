package router 

import  "social_network/internal/api/v1"

func init() {
	APIRouter.HandleFunc("/home", v1.Authentication(v1.Home)).Methods("GET", "POST")	
}
