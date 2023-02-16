package router

import (
	"fmt"
	"net/http"
	"social_network/internal/api/router/options"
	"social_network/internal/api/v1"
	"social_network/internal/socket"
)

func init() {
	router.APIRouter.HandleFunc("/home", v1.Authentication(v1.Home))

	hub := socket.NewHub()
	go hub.Run()

	router.APIRouter.HandleFunc("/ws", v1.Authentication(func(wrt http.ResponseWriter, req *http.Request) {
		id := req.URL.Query().Get("id")
		client := socket.NewClient(id, hub, wrt, req)
		client.Hub.Register <- client
		fmt.Println("Client:", client.ID)
	}))
}
