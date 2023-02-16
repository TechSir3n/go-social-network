package app

import (
	"context"
	"net/http"
	"os"
	"time"

	_ "social_network/internal/api/router"
	_ "social_network/internal/api/router/github"
	_ "social_network/internal/api/router/google"
	"social_network/internal/api/router/options"
	_ "social_network/internal/api/router/request"
	_ "social_network/internal/api/router/static"
	sql "social_network/internal/config/database"
	"social_network/utils/logger"
	"social_network/utils/password"
)

func Start() {
	sql.ConnectDB()

	defer sql.ConnectDB().Close(context.Background())
}

func init() {
	generator.GeneratePassword()

	srv := &http.Server{
		Handler:      router.APIRouter,
		Addr:         ":" + os.Getenv("PORT"),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	logger.Fatal(srv.ListenAndServe())
}
