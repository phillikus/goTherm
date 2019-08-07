package api

import (
	"encoding/json"
	"github.com/phillikus/goTherm/src/webapp/api/interfaces"
	"github.com/phillikus/goTherm/src/webapp/config"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func HandlePushThermData(config *config.Config, repository api_interfaces.IThermRepository) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			add(responseWriter, request, repository)
		} else if request.Method == "GET" {
			getLatest(responseWriter, request, repository)
		}
	}
}

func getLatest(responseWriter http.ResponseWriter, request *http.Request, repository api_interfaces.IThermRepository) {
	lastEntries := repository.GetLatest(1)
	json, err := json.Marshal(lastEntries)

	if err != nil {
		panic(err)
	}

	responseWriter.Write(json)
}

func add(responseWriter http.ResponseWriter, request *http.Request, repository api_interfaces.IThermRepository) {
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	if err != nil {
		panic(err)
	}

	message := string(body)
	temp, err := strconv.ParseFloat(message, 32)

	if err == nil {
		repository.Insert(api_interfaces.ThermData{
			CreateDate:  time.Now(),
			Temperature: float32(temp),
		})
	} else {
		panic(err)
	}
}
