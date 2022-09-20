package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStoringTakenDosageAndRetrievingIt(t *testing.T) {
	store := NewInMemoryDataStore()
	server := supplementsHandler{store}
	supplement := "Magnesium"
	/*
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	*/

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetSupplementDosage(supplement))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "123")
}
