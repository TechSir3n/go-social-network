package v1

import (
	"net/http"
	"social_network/utils"
)


func Home(wrt http.ResponseWriter, req *http.Request){
	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/home/main.html", nil)
}