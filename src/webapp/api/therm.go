package api

import (
	"fmt"
	"github.com/phillikus/goTherm/src/webapp/config"
	"io/ioutil"
	"net/http"
)

func HandlePushThermData(config *config.Config) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			pushThermData(responseWriter, request)
		}
	}
}

func pushThermData(responseWriter http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	message := string(body)

	if err == nil {
		fmt.Fprintf(responseWriter, message)
	} else {
		panic("An error has ocurred")
	}
}
