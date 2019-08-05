package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandlePushThermData() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		pushThermData(responseWriter, request)
	}
}

func pushThermData(responseWriter http.ResponseWriter, request *http.Request) {
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
