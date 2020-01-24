package api

import (
	"bytes"
	api "github.com/goTherm/api/interfaces"
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
	handler := HandlePushThermData(nil, &thermRepositoryFake{})

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

type thermRepositoryFake struct {
}

func (repository *thermRepositoryFake) Insert(data api.ThermData) {
	return
}

func (repository *thermRepositoryFake) GetLatest(count int) []api.ThermData {
	return []api.ThermData{}
}
