package main

import (
	"fmt"
	"github.com/phillikus/goTherm/src/webapp/api"
	"log"
	"net/http"
	"os"
)

func home(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "static/index.html")
}

func handleLogging(next http.HandlerFunc, logger *log.Logger) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		logger.Output(1, fmt.Sprintf("Begin request: %s", request.RequestURI))
		next(responseWriter, request)
		logger.Output(1, fmt.Sprintf("End request: %s", request.RequestURI))
	}
}

func handleRequests() {
	fileServer := http.FileServer(http.Dir("static"))
	logger := log.New(os.Stderr, "", log.LstdFlags)

	http.HandleFunc("/", handleLogging(home, logger))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/therm", handleLogging(api.HandlePushThermData(), logger))

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		log.Fatal("An error ocurred")
	}
}

func main() {
	handleRequests()
}
