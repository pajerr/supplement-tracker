package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingTakenDosagesAndRetrievingThem(t *testing.T) {
	store := NewInMemorySupplementStore()
	server := NewSupplementsServer(store)

	supplement := "magnesium"

	//store taken magensium units 3 times with POST request
	server.ServeHTTP(httptest.NewRecorder(), newPostUnitsTakenRequest(supplement))
	server.ServeHTTP(httptest.NewRecorder(), newPostUnitsTakenRequest(supplement))
	server.ServeHTTP(httptest.NewRecorder(), newPostUnitsTakenRequest(supplement))

	t.Run("get taken supplement units", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetUnitsTakenRequest(supplement))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	supplement = "vitamin-c"
	server.ServeHTTP(httptest.NewRecorder(), newPostUnitsTakenRequest(supplement))

	t.Run("get All Supplements status", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetDashboardRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getDashboardFromResponse(t, response.Body)
		want := []Supplement{
			{"magnesium", 3},
			{"vitamin-c", 1},
		}
		assertGetDashboard(t, got, want)
	})
}

/*
//need to implement set/change the taken supplement dose in the data store
func TestRetieveSupplementDosages(t *testing.T) {
	store := NewInMemorySupplementStore()
	server := NewSupplementsServer(store)

	supplement := "magnesium"

	//set the dosage for magnesium, dose is hardcoded to 400
	server.ServeHTTP(httptest.NewRecorder(), newSetSupplementDosage(supplement))
	response := httptest.NewRecorder()

	server.ServeHTTP(httptest.NewRecorder(), newGetSupplementDosage(supplement))

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "400")
}
*/
