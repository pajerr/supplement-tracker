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

//function to save taken supplement dosage
func (stub *StubSupplementDataStore) StoreTakenDosage(name string, dosage int) {
	stub.dosages[name] = dosage
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
	})
}

//Testing that POST reponse gets accepted
func TestPostAccepted(t *testing.T) {
	store := StubSupplementDataStore{
		map[string]int{},
	}

	server := &supplementsHandler{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/supplements/magnesium", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})
}

func TestStoringTakenDosage(t *testing.T) {
	//create stub data store
	store := StubSupplementDataStore{
		map[string]int{
			"vitamin-c": 0,
			"magnesium": 0,
		},
	}
	server := &supplementsHandler{&store}

	t.Run("it stores taken Magnesium dosage when POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/supplements/magnesium/200/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		//assertStatus(t, response.Code, http.StatusAccepted)

		if store.dosages["magnesium"] != 300 {
			t.Errorf("got %d, want %d", store.dosages["magnesium"], 300)
		}
		//assertResponseBody(t, response.Body.String(), "200")
	})
}

//Helper functon to create a new GET request for a supplement
func newGetSupplementDosage(supplementName string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/supplements/%s", supplementName), nil)
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
