package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//A map is a quick and easy way of making a stub key/value store for our tests
//dosages stores name and dosage of single supplement
//taken dosges shows daily taken amount of those dosages and it can be combined to total daily dosage
type StubSupplementDataStore struct {
	dosages      map[string]int
	takenDosages []string
}

func (stub *StubSupplementDataStore) GetSupplementDosage(name string) int {
	dosage := stub.dosages[name]
	return dosage
}

//record the taken dosage
func (stub *StubSupplementDataStore) RecordTakenDosage(name string) {
	stub.takenDosages = append(stub.takenDosages, name)
}

func TestTakenSupplementDosage(t *testing.T) {
	//create stub data store
	store := StubSupplementDataStore{
		map[string]int{
			"vitamin-c": 500,
			"magnesium": 400,
		},
		nil,
	}

	//create a new instance of our supplementsHandler and then call its method ServeHTTP
	//send in the stub data store as the argument to the supplementsHandler/server
	server := &supplementsServer{&store}

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
	})
}

//Testing that POST reponse gets accepted, only tests that the status code is 200
func TestStoreTakenDosage(t *testing.T) {
	store := StubSupplementDataStore{
		map[string]int{},
		nil,
	}

	server := &supplementsServer{&store}

	t.Run("it records taken dosage when POST", func(t *testing.T) {
		supplement := "magnesium"
		request := newPostTakenDosageRequest(supplement)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.takenDosages) != 1 {
			t.Errorf("got %d want %d", len(store.takenDosages), 1)
		}

		if store.takenDosages[0] != supplement {
			t.Errorf("did not store correct supplement got %q want %q", store.takenDosages[0], supplement)
		}

	})
}

//Helper functon to create a new GET request for a supplement
func newGetSupplementDosage(supplementName string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/supplements/%s", supplementName), nil)
	return req
}

//Helper function to create a new POST request for a taken daily dosage for supplement
func newPostTakenDosageRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/supplements/%s", name), nil)
	return req
}

//Helper function to check the response body
func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

//Helper function to check the response status
func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
