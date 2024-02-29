package main

import (
	"log"

	"github.com/samallen659/ccWebServer/internal/server"
)

func main() {
	svr, err := server.NewServer("localhost:8080")
	if err != nil {
		log.Panic(err)
	}

	if err := svr.Listen(); err != nil {
		log.Panic(err)
	}
}
