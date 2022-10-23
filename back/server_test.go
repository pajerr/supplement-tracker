package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

//A map is a quick and easy way of making a stub key/value store for our tests
//dosages stores name and dosage of single supplement
//taken dosges shows daily taken amount of those dosages and it can be combined to total daily dosage
type StubSupplementDataStore struct {
	//supplement name and dosage
	dosages map[string]int
	//dose of single unit of supplement
	//UnitDose map[string]int
	//supplement name and taken units
	UnitsTaken map[string]int
	//lists all taken supplements and taken units
	Dashboard []Supplement
}

// ##### /dosage functions #####
func (stub *StubSupplementDataStore) GetSupplementDosage(name string) int {
	dosage := stub.dosages[name]
	return dosage
}

func (stub *StubSupplementDataStore) SetSupplementDosage(name string, dosage int) {
	stub.dosages[name] = dosage
}

// ##### /supplement function ######
func (stub *StubSupplementDataStore) GetUnitsTaken(name string) int {
	UnitsTaken := stub.UnitsTaken[name]
	return UnitsTaken
}

//record the taken supplement dose
func (stub *StubSupplementDataStore) RecordUnitsTaken(name string) {
	stub.UnitsTaken[name]++
}

// #### /dashboard functions ###
func (s *StubSupplementDataStore) GetDashboard() []Supplement {
	return s.Dashboard
}

func TestUnitsTakenDosage(t *testing.T) {
	//create stub data store
	store := StubSupplementDataStore{
		map[string]int{
			"vitamin-c": 500,
			"magnesium": 400,
		},
		map[string]int{},
		//pass nil for Dashboard
		nil,
	}

	//create a new instance of our supplementsHandler and then call its method ServeHTTP
	//send in the stub data store as the argument to the supplementsHandler/server
	server := NewSupplementsServer(&store)

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
func TestStoreUnitsTaken(t *testing.T) {
	store := StubSupplementDataStore{
		map[string]int{},
		map[string]int{},
		//pass nil for Dashboard
		nil,
	}

	server := NewSupplementsServer(&store)

	t.Run("it records taken supplemenet when POST", func(t *testing.T) {
		supplement := "magnesium"
		request := newPostUnitsTakenRequest(supplement)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		//check that length of taken supplements is 1
		if len(store.UnitsTaken) != 1 {
			t.Errorf("got %d want %d", len(store.UnitsTaken), 1)
		}

		//store second taken supplement dose
		server.ServeHTTP(response, request)

		//check that amount of taken dosages for supplement is 2 as expected, since we made 2 POST requests
		if store.UnitsTaken[supplement] != 2 {
			t.Errorf("got %d want %d", store.UnitsTaken[supplement], 2)
		}

		/*
			if store.UnitsTaken[0] != supplement {
				t.Errorf("did not store correct supplement got %q want %q", store.UnitsTaken[0], supplement)
			}
		*/

	})
}

//server_test.go
func TestListAllTakenSupps(t *testing.T) {

	t.Run("it returns supplementStatus as JSON", func(t *testing.T) {
		wantedDashboard := []Supplement{
			{"vitamin-c", 1},
			{"magnesium", 2},
			{"iron", 1},
		}

		store := StubSupplementDataStore{nil, nil, wantedDashboard}
		server := NewSupplementsServer(&store)

		request := newGetDashboardRequest()
		//request, _ := http.NewRequest(http.MethodGet, "/dashboard", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		//server_test.go
		if response.Result().Header.Get("content-type") != "application/json" {
			t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
		}

		//helper function handles error checking also
		got := getDashboardFromResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)
		assertGetDashboard(t, got, wantedDashboard)
	})
}

//Helper functon to create a new GET request for a supplement dosage
func newGetSupplementDosage(supplementName string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/dosages/%s", supplementName), nil)
	return req
}

func newSetSupplementDosage(supplementName string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/dosages/%s", supplementName), nil)
	return req
}

func newGetUnitsTakenRequest(supplementName string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/supplements/%s", supplementName), nil)
	return req
}

//Helper function to create a new POST request for a taken daily dosage for supplement
func newPostUnitsTakenRequest(name string) *http.Request {
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

//helper to check header content type
func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

//dashboard path helpers
func getDashboardFromResponse(t testing.TB, body io.Reader) (Dashboard []Supplement) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&Dashboard)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Supplement, '%v'", body, err)
	}

	return
}

//dashboard path helper for new http requests
func newGetDashboardRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/dashboard", nil)
	return req
}

func assertGetDashboard(t testing.TB, got, want []Supplement) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
