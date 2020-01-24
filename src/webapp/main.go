package main

import (
	"fmt"
	"os"

	api "github.com/goTherm/api"
	config "github.com/goTherm/config"
)

func main() {
	fmt.Println("Starting up API!")
	args := os.Args[1:]
	var cfg config.Config

	if len(args) > 0 && args[0] == "-prod" {
		cfg = config.LoadConfigurationFromEnv()
	} else {
		cfg = config.LoadConfigurationFromFile("config.json")
	}

	api.InitRoutes(&cfg)
}
