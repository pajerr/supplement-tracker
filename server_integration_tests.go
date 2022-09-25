//server_integration_test.go

package main

/*
func TestRecordingTakenDosagesAndRetrievingThem(t *testing.T) {
	store := InMemorySupplementDataStore{}
	server := supplementsServer{&store}
	supplement := "magnesium"

	server.ServeHTTP(httptest.NewRecorder(), newPostTakenDosageRequest(supplement))
	server.ServeHTTP(httptest.NewRecorder(), newPostTakenDosageRequest(supplement))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "2")
}
*/
