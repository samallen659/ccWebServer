package main

import (
	"fmt"
	"log"
	"os"

	"github.com/samallen659/ccWebServer/internal/configuration"
	"github.com/samallen659/ccWebServer/internal/server"
	"gopkg.in/yaml.v2"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Panicf("Failed accessing working dir: %s", err.Error())
	}

	cb, err := os.ReadFile(fmt.Sprintf("%s/config.yaml", wd))
	if err != nil {
		log.Panicf("Failed reading config file: %s", err.Error())
	}

	var config configuration.Config
	err = yaml.Unmarshal(cb, &config)
	if err != nil {
		log.Panicf("Failed Unmarshalling yaml file into config struct: %s", err.Error())
	}

	svr, err := server.NewServer(config.ListenAddr, config.WWWPath)
	if err != nil {
		log.Panicf("Failed creating server: %s", err.Error())
	}

	if err := svr.Listen(); err != nil {
		log.Panic(err)
	}
}
