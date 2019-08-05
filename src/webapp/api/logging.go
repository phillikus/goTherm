package api

import (
	"fmt"
	"log"
	"net/http"
)

func HandleLogging(next http.HandlerFunc, logger *log.Logger) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		logger.Output(1, fmt.Sprintf("Begin request: %s", request.RequestURI))

		if next != nil {
			next(responseWriter, request)
		}

		logger.Output(1, fmt.Sprintf("End request: %s", request.RequestURI))
	}
}
