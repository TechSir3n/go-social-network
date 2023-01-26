package v2

import (
	"net/http"

	"github.com/gorilla/mux"
	"social_network/internal/api/v1"
)

func hello(w http.ResponseWriter, r *http.Request) {

}

func InitRouter() *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/registration", v1.SignUp).Methods("GET","POST")

	mux.HandleFunc("/login", v1.Login).Methods("GET","POST")

	mux.HandleFunc("/home", hello).Methods("GET","POST")

	mux.HandleFunc("/reset/password", hello).Methods("GET","POST")

	return mux
}
