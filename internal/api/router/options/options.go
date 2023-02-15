package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

var APIRouter = mux.NewRouter().Schemes("http").Subrouter().StrictSlash(true)

func init() {
	APIRouter.Methods("OPTIONS").HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		wrt.Header().Set("Access-Control-Allow-Origin", "")
		wrt.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT,DELETE")
		wrt.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})

}
