package main

import (
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

	logger.Info("register user handler")
	handler := user.NewHandler(*logger)
	handler.Register(router)

	startServer(router)
}

func startServer(router *httprouter.Router) {
	logger := log.GetLogger()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	logger.Info("listening on port 8888")

	logger.Fatal(server.Serve(listener))
}
