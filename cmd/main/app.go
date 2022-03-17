package main

import (
	"fmt"
	"go-rest/internal/config"
	"go-rest/internal/user"
	"go-rest/pkg/log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := log.GetLogger()
	logger.Info("initialized router")
	router := httprouter.New()
	conf := config.GetConfig()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	startServer(router, conf)
}

func startServer(router *httprouter.Router, conf *config.Config) {
	logger := log.GetLogger()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", conf.Host, conf.Port))
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	logger.Infof("listening on port %s", conf.Port)

	logger.Fatal(server.Serve(listener))
}
