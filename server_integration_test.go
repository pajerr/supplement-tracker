//server_integration_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingTakenDosagesAndRetrievingThem(t *testing.T) {
	store := InMemorySupplementDataStore{}
	server := supplementsServer{&store}
	supplement := "magnesium"

	server.ServeHTTP(httptest.NewRecorder(), newPostTakenSupplementRequest(supplement))
	server.ServeHTTP(httptest.NewRecorder(), newPostTakenSupplementRequest(supplement))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetTakenSupplementRequest(supplement))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "2")
}
