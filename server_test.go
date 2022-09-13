package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//A map is a quick and easy way of making a stub key/value store for our tests
type StubSupplementDataStore struct {
	dosages map[string]int
}

func (stub *StubSupplementDataStore) GetSupplementDosage(name string) int {
	dosage := stub.dosages[name]
	return dosage
}

func TestSupplementDosages(t *testing.T) {
	//create stub data store
	store := StubSupplementDataStore{
		map[string]int{
			"vitamin-c": 500,
			"magnesium": 400,
		},
	}

	//create a new instance of our supplementsHandler and then call its method ServeHTTP
	//send in the stub data store as the argument to the supplementsHandler/server
	server := &supplementsHandler{&store}

	t.Run("Return Vitamin C dosage", func(t *testing.T) {
		//Use helper function to create a new GET request for Vitamin C
		request := newGetSupplementDosage("vitamin-c")
		response := httptest.NewRecorder()

		//we pass in the response and request to the ServeHTTP method from our supplementsHandler
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "500")
	})

	t.Run("Return Magnesium dosage", func(t *testing.T) {
		//Use helper function to create a new GET request for Vitamin C
		request := newGetSupplementDosage("magnesium")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "400")
	})

	t.Run("Return 404 on missing supplement", func(t *testing.T) {
		request := newGetSupplementDosage("iron")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
		assertResponseBody(t, response.Body.String(), "Supplement not found")
	})
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

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
