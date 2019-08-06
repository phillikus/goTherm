package main

import (
	"github.com/phillikus/goTherm/src/webapp/api"
	"github.com/phillikus/goTherm/src/webapp/config"
)

func main() {
	config := config.LoadConfiguration("config.json")
	api.InitRoutes(&config)
}
