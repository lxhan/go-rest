package main

import (
	"go-rest/internal/user"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	handler := user.NewHandler()
	handler.Register(router)

	startServer(router)
}

func startServer(router *httprouter.Router) {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	server.Serve(listener)
	log.Println("Server is listening")
}
