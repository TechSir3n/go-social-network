package router

import (
	"social_network/internal/api/v1"
	"social_network/internal/api/router/options"
)

func init() {
	router.APIRouter.HandleFunc("/admin", v1.Authentication(v1.AdminIndex))
}
