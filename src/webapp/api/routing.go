package api

import (
	"github.com/phillikus/goTherm/src/webapp/config"
	"log"
	"net/http"
	"os"
)

func InitRoutes(config *config.Config) {
	fileServer := http.FileServer(http.Dir("static"))
	logger := log.New(os.Stderr, "", log.LstdFlags)

	http.HandleFunc("/", HandleLogging(home, logger))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/therm", HandleLogging(HandlePushThermData(config), logger))

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		logger.Fatal("An error ocurred")
	}
}

func home(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "static/index.html")
}
