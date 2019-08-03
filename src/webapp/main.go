package main

import (
	"github.com/phillikus/goTherm/src/webapp/api"
	"log"
	"net/http"
)

func home(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "static/index.html")
}

func handleRequests() {
	fileServer := http.FileServer(http.Dir("static"))
	http.HandleFunc("/", home)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/therm", api.PushThermData)
	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		log.Fatal("An error ocurred")
	}
}

func main() {
	handleRequests()
}

//http://192.168.0.35:5000/therm/push
