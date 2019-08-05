package api

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/test123", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "text/plain")
	req.RequestURI = "/test123"

	var logStream bytes.Buffer
	logger := log.New(&logStream, "", 0)

	requestRecorder := httptest.NewRecorder()
	handler := HandleLogging(nil, logger)

	handler.ServeHTTP(requestRecorder, req)

	expected := "Begin request: /test123\nEnd request: /test123\n"
	actual := logStream.String()

	if expected != actual {
		t.Errorf("Wrong log message, expected: %v but was %v", expected, actual)
	}
}
