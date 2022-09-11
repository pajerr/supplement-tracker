package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSupplementDosages(t *testing.T) {
	//create a new instance of our supplementsHandler and then call its method ServeHTTP
	server := &supplementsHandler{}

	t.Run("Return Vitamin C dosage", func(t *testing.T) {
		//Use helper function to create a new GET request for Vitamin C
		request := newGetSupplementDosage("vitamin-c")
		response := httptest.NewRecorder()

		//we pass in the response and request to the ServeHTTP method from our supplementsHandler
		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "500")
	})

	t.Run("Return Magnesium dosage", func(t *testing.T) {
		//Use helper function to create a new GET request for Vitamin C
		request := newGetSupplementDosage("magnesium")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "400")
	})

	/*t.Run("Return 404 on missing supplement", func(t *testing.T) {
		//Use helper function to create a new GET request for Vitamin C
		request := newGetSupplementDosage("iron")
		response := httptest.NewRecorder()

		supplementsHandler(response, request)
		assertStatus(t, response.Code, http.StatusNotFound)
		assertResponseBody(t, response.Body.String(), "Supplement not found")
	})*/
}

func newGetSupplementDosage(supplementName string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/supplements/%s", supplementName), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
