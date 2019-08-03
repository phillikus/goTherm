package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PushThermData(responseWriter http.ResponseWriter, request *http.Request) {
	log.Output(1, "Calling /therm...")

	if request.Method != "POST" {
		panic("Only POST allowed for this method")
	}

	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	message := string(body)

	if err == nil {
		fmt.Fprintf(responseWriter, message)
	} else {
		panic("An error has ocurred")
	}
}
