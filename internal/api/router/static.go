package router

import "net/http"

func init() {
	// i used absolute path because the other one didn't work
	access_css := http.FileServer(http.Dir("C:/Users/Ruslan/Desktop/go-social-network/static/access/css/"))
	access_js := http.FileServer(http.Dir("C:/Users/Ruslan/Desktop/go-social-network/static/access/js/"))
	home_css := http.FileServer(http.Dir("C:/Users/Ruslan/Desktop/go-social-network/static/home/css1/"))
	home_js := http.FileServer(http.Dir("C:/Users/Ruslan/Desktop/go-social-network/static/home/js1/"))
	admin_css := http.FileServer(http.Dir("C:/Users/Ruslan/Desktop/go-social-network/static/admin/css2/"))
	admin_js := http.FileServer(http.Dir("C:/Users/Ruslan/Desktop/go-social-network/static/admin/js2/"))

	APIRouter.PathPrefix("/css/").Handler(http.StripPrefix("/css/", access_css))

	APIRouter.PathPrefix("/js/").Handler(http.StripPrefix("/js/", access_js))

	APIRouter.PathPrefix("/css1/").Handler(http.StripPrefix("/css1/", home_css))

	APIRouter.PathPrefix("/js1/").Handler(http.StripPrefix("/js1/", home_js))

	APIRouter.PathPrefix("/css2/").Handler(http.StripPrefix("/css2/", admin_css))

	APIRouter.PathPrefix("/js2/").Handler(http.StripPrefix("/js2/", admin_js))
}
