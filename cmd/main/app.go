package main

import (
	"fmt"
	"go-rest/internal/config"
	"go-rest/internal/user"
	"go-rest/pkg/logger"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log := logger.GetLogger()
	log.Info("initialized router")
	router := httprouter.New()
	conf := config.GetConfig()

	log.Info("register user handler")
	handler := user.NewHandler(log)
	handler.Register(router)

	startServer(router, conf)
}

func startServer(router *httprouter.Router, conf *config.Config) {
	log := logger.GetLogger()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", conf.Host, conf.Port))
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Infof("listening on port %s", conf.Port)

	log.Fatal(server.Serve(listener))
}
