package api

import (
	"encoding/json"
	api "github.com/phillikus/goTherm/api/interfaces"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	config "github.com/phillikus/goTherm/config"
)

func HandlePushThermData(config *config.Config, repository api.ThermRepository) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			add(responseWriter, request, repository)
		} else if request.Method == "GET" {
			getLatest(responseWriter, request, repository)
		}
	}
}

func getLatest(responseWriter http.ResponseWriter, request *http.Request, repository api.ThermRepository) {
	lastEntries := repository.GetLatest(1)
	json, err := json.Marshal(lastEntries)

	if err != nil {
		panic(err)
	}

	responseWriter.Write(json)
}

func add(responseWriter http.ResponseWriter, request *http.Request, repository api.ThermRepository) {
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	if err != nil {
		panic(err)
	}

	message := string(body)
	temp, err := strconv.ParseFloat(message, 32)

	if err == nil {
		repository.Insert(api.ThermData{
			CreateDate:  time.Now(),
			Temperature: float32(temp),
		})
	} else {
		panic(err)
	}
}
