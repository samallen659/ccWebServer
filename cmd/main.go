package main

import (
	"fmt"
	"log"
	"os"

	"github.com/samallen659/ccWebServer/internal/server"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	wwwPath := fmt.Sprintf("%s/www", root)
	svr, err := server.NewServer("localhost:8080", wwwPath)
	if err != nil {
		log.Panic(err)
	}

	if err := svr.Listen(); err != nil {
		log.Panic(err)
	}
}
