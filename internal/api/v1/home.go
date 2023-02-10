package v1

import (
	"net/http"
	"social_network/utils"
	_ "social_network/utils/logger"
)

func Home(wrt http.ResponseWriter, req *http.Request) {

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/home/html/main.html", nil)
}

func Message(wrt http.ResponseWriter, req *http.Request) {

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/home/html/message.html", nil)
}

func Friends(wrt http.ResponseWriter, req *http.Request) {

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/home/html/friends.html", nil)
}

func Music(wrt http.ResponseWriter, req *http.Request) {

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/home/html/music.html", nil)
}

func Video(wrt http.ResponseWriter, req *http.Request) {

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/home/html/video.html", nil)
}

func Bookmarks(wrt http.ResponseWriter, req *http.Request) {

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/home/html/bookmarks.html", nil)
}

func Settings(wrt http.ResponseWriter, req *http.Request) {

	utils.ExecTemplate(wrt, "C:/Users/Ruslan/Desktop/go-social-network/static/home/html/settings.html", nil)
}
