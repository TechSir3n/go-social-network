package utils

import (
	templ "html/template"
	"log"
	"net/http"
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
