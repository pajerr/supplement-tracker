package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingTakenDosagesAndRetrievingThem(t *testing.T) {
	store := NewInMemorySupplementStore()
	//server := NewSupplementsServer{store}
	server := NewSupplementsServer(store)

	//store := InMemorySupplementDataStore{}
	//server := supplementsServer{&store}
	supplement := "magnesium"

	server.ServeHTTP(httptest.NewRecorder(), newPostTakenSupplementRequest(supplement))
	server.ServeHTTP(httptest.NewRecorder(), newPostTakenSupplementRequest(supplement))
	server.ServeHTTP(httptest.NewRecorder(), newPostTakenSupplementRequest(supplement))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetTakenSupplementRequest(supplement))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

//need to implement set/change the taken supplement dose in the data store
func TestRetieveSupplementDosages(t *testing.T) {
	store := NewInMemorySupplementStore()
	server := NewSupplementsServer(store)

	supplement := "magnesium"

	server.ServeHTTP(httptest.NewRecorder(), newGetSupplementDosage(supplement))
	response := httptest.NewRecorder()

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "400")
}
