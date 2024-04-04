package main

import (
	"app/api"
	"log"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	if err := server.Start(":3000"); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
