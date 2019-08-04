package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPushThermData(t *testing.T) {
	message := []byte(`20.00`)
	req, err := http.NewRequest("POST", "/therm", bytes.NewBuffer(message))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "text/plain")

	requestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(PushThermData)

	handler.ServeHTTP(requestRecorder, req)

	if status := requestRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v instead of %v", status, http.StatusOK)
	}
}

func TestPushThermDataInvalidMethod(t *testing.T) {
	_, err := http.NewRequest("GET", "/therm/push", nil)

	if err != nil {
		panic("Invalid request type should panic")
	}
}
