package app

import (
	 "context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "social_network/internal/api/router"
	_ "social_network/internal/api/router/github"
	_ "social_network/internal/api/router/google"
	"social_network/internal/api/router/options"
	_ "social_network/internal/api/router/request"
	_ "social_network/internal/api/router/static"
	"social_network/internal/api/v1/middleware"
	sql  "social_network/internal/config/database"
	"social_network/utils/logger"
	"social_network/utils/password"
)

func Start() {
	sql.ConnectDB()

	defer sql.ConnectDB().Close(context.Background())

	HTTP := flag.Bool("http", true, "run server http")
	key := flag.String("key", "go-server.key", "key PEM file")
	crt := flag.String("crt", "go-server.crt", "certificate PEM file")
	flag.Parse()

	fmt.Println("Http:", *HTTP)

	generator.GeneratePassword()

	srv := &http.Server{
		Handler:      middleware.Logging(router.APIRouter),
		Addr:         ":" + os.Getenv("PORT"),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	for arg := range flag.Args() {
		fmt.Println("No Command line: ", arg)
	}
	
	//http or https 
	if *HTTP {
		logger.Fatal(srv.ListenAndServe())
	} else {
		logger.Fatal(srv.ListenAndServeTLS(*crt,*key))
	}
}
