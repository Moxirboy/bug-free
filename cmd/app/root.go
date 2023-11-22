package app

import (
	http2 "bug-free/internal/http"
	"bug-free/internal/server"
	"log"
)




// @title Your API Title
// @version 1.0
// @description Your API Description
// @host localhost:8080
// @BasePath /api/v1

func Execute() {

	srv := new(server.Server)
	port := "3000"
	log.Print("started")
	if err := srv.Run(port, http2.Router()); err != nil {
		log.Fatalf("error occured")
	}
}
