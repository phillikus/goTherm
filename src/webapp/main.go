package main

import (
	"github.com/phillikus/goTherm/src/webapp/api"
	"github.com/phillikus/goTherm/src/webapp/config"
	"os"
)

func main() {
	args := os.Args[1:]
	var cfg config.Config

	if len(args) > 0 && args[0] == "-prod" {
		cfg = config.LoadConfigurationFromEnv()
	} else {
		cfg = config.LoadConfigurationFromFile("config.json")
	}

	api.InitRoutes(&cfg)
}
