package router

import v1 "social_network/internal/api/v1"
import "net/http"

func init() {
	APIRouter.HandleFunc("/user", v1.Authentication(v1.UserIndex)).Methods("GET")

	APIRouter.Methods("OPTIONS").HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		wrt.Header().Set("Access-Control-Allow-Origin", " ")
		wrt.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		wrt.Header().Set("Access-Control-Allow-Headers", "Content-Language, Content-Type, Cache-Control, Content-Length, Authorization")
	})
}
