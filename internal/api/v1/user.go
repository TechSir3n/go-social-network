package v1

import "net/http"

func UserIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome, User"))
}

func AdminIndex(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome Admin"))
}