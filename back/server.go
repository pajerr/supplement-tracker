package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//Interface of datastore, functions are defined in main.go
type SupplementDataStore interface {
	SetSupplementDosage(supplement string, dosage int)
	GetSupplementDosage(name string) int
	RecordTakenSupplement(name string)
	GetTakenSupplement(name string) int
	//get status for all supplements from /listtaken endpoint
	GetAllSupplementsStatus() []Supplement
}

//type for /listtaken
type Supplement struct {
	Name         string
	DosagesTaken int
}

//consant to return correct header content-type
const jsonContentType = "application/json"

//allows us to use the SupplementDataStore interface in the handler
//for example to store.GetSupplelementDosage to get supplements dosage
//supplementsServer now has all the methods that http.Handler has, which is just ServeHTTP, this is embedding
//When embedding types, really think about what impact that has on your public API.
//It is a very common mistake to misuse embedding and end up polluting your APIs and exposing the internals of your type.
type supplementsServer struct {
	store SupplementDataStore
	//router *http.ServeMux
	http.Handler
}

//do the one-time setup of creating the router. Each request can then just use that one instance of the router
func NewSupplementsServer(store SupplementDataStore) *supplementsServer {
	/*s := &supplementsServer{
		store,
		http.NewServeMux(),
	}*/

	s := new(supplementsServer)

	s.store = store

	router := http.NewServeMux()
	router.Handle("/dosages/", http.HandlerFunc(s.dosagesHandler))
	router.Handle("/supplements/", http.HandlerFunc(s.supplementsHandler))
	router.Handle("/listtaken", http.HandlerFunc(s.listTakenSupplementsHandler))

	s.Handler = router

	return s
}

/*
//Refactored to use SupplementDataStore interface
func (s *supplementsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
*/

// ############################### /listtaken path ###############################

func (s *supplementsServer) listTakenSupplementsHandler(w http.ResponseWriter, r *http.Request) {
	//old hardocded version
	/*supplementTakenTable := []Supplement{
		{"Magnesium", 2},
	}*/

	//Set header from const jsonContentType
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(s.store.GetAllSupplementsStatus())
	//json.NewEncoder(w).Encode(supplementTakenTable)

	w.WriteHeader(http.StatusOK)
}

// ############################### /dosages path ###############################

func (s *supplementsServer) dosagesHandler(w http.ResponseWriter, r *http.Request) {
	supplement := strings.TrimPrefix(r.URL.Path, "/dosages/")

	switch r.Method {
	//not yet implemented
	case http.MethodPost:
		s.processSetDosage(w, supplement)
	case http.MethodGet:
		s.showDosage(w, supplement)
	}
}

// ############################### /supplements path ###############################

func (s *supplementsServer) supplementsHandler(w http.ResponseWriter, r *http.Request) {
	supplement := strings.TrimPrefix(r.URL.Path, "/supplements/")

	switch r.Method {
	case http.MethodPost:
		s.processTakenSupplement(w, supplement)
	case http.MethodGet:
		s.showTakenSupplement(w, supplement)
	}
}

// ############################### /dosages path functions ###############################

// dosages GET function
func (s *supplementsServer) showDosage(w http.ResponseWriter, supplement string) {
	//r.URL.Path returns the path of the request which we can then use strings.TrimPrefix to trim away /supplements/
	dosage := s.store.GetSupplementDosage(supplement)

	//if supplement not in StubSupplementDataStore return 404
	if supplement != "magnesium" && supplement != "vitamin-c" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, dosage)

}

// dosages POST function
func (s *supplementsServer) processSetDosage(w http.ResponseWriter, supplement string) {
	s.store.SetSupplementDosage(supplement, 400)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
}

// ############################### /supplements path functions ###############################

// /supplements POST function
// Function to process the POST request on /supplements/supplement handler and record 1 dosage as taken
func (s *supplementsServer) processTakenSupplement(w http.ResponseWriter, supplement string) {
	//return 200 status code if request is POST method to pass test
	s.store.RecordTakenSupplement(supplement)

	//to fix CORS errors in frontend
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
}

// /supplements GET function
func (s *supplementsServer) showTakenSupplement(w http.ResponseWriter, supplement string) {
	takenSuppAmount := s.store.GetTakenSupplement(supplement)

	if takenSuppAmount == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	//to fix CORS errors in frontend
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, takenSuppAmount)
}
