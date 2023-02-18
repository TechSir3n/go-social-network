package utils

import (
	templ "html/template"
	"log"
	"math/rand"
	"net/http"
)

const (
	Path = "C:/Users/Ruslan/Desktop/go-social-network/static"
)

func ExecTemplate(w http.ResponseWriter, template string, data interface{}) {
	t, err := templ.ParseFiles(template)
	if err != nil {
		log.Fatal("Failed to read html file")
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// this is a random string needed to set in oauth state
func GenerateRandomString() string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, 5)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}

	return string(str)
}
