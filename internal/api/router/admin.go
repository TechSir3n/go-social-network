package router

import "social_network/internal/api/v1"

func init() {
	APIRouter.HandleFunc("/admin", v1.Authentication(v1.AdminIndex)).Methods("GET", "POST")

	APIRouter.HandleFunc("/admin/{method_find}/find", v1.Authentication(v1.AdminIndex)).Methods("GET", "POST")

	APIRouter.HandleFunc("/admin/{method_delete}/delete", v1.Authentication(v1.AdminIndex)).Methods("GET", "DELETE")

	APIRouter.HandleFunc("/admin/{method_get}/get", v1.Authentication(v1.AdminIndex)).Methods("GET", "POST")

	APIRouter.HandleFunc("/admin/{method_get_all}/get_all", v1.Authentication(v1.AdminIndex)).Methods("GET", "POST")
}
