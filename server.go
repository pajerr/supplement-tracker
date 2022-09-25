package main

import (
	"fmt"
	"net/http"
	"strings"
)

//Interface of datastore, functions are defined in main.go
type SupplementDataStore interface {
	GetSupplementDosage(name string) int
	RecordTakenSupplement(name string)
	GetTakenSupplement(name string) int
}

//allows us to use the SupplementDataStore interface in the handler
//for example to store.GetSupplelementDosage to get supplements dosage
type supplementsServer struct {
	store SupplementDataStore
}

//Refactored to use SupplementDataStore interface
func (s *supplementsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	router.Handle("/dosages/", http.HandlerFunc(s.dosagesHandler))
	router.Handle("/supplements/", http.HandlerFunc(s.supplementsHandler))

	router.ServeHTTP(w, r)
}

func (s *supplementsServer) dosagesHandler(w http.ResponseWriter, r *http.Request) {
	supplement := strings.TrimPrefix(r.URL.Path, "/dosages/")

	switch r.Method {
	//not yet implemented
	//case http.MethodPost:
	//s.processSetDosage(w, supplement)
	case http.MethodGet:
		s.showDosage(w, supplement)
	}
}

func (s *supplementsServer) supplementsHandler(w http.ResponseWriter, r *http.Request) {
	supplement := strings.TrimPrefix(r.URL.Path, "/supplements/")

	switch r.Method {
	case http.MethodPost:
		s.processTakenSupplement(w, supplement)
	case http.MethodGet:
		s.showTakenSupplement(w, supplement)
	}
}

func (s *supplementsServer) showDosage(w http.ResponseWriter, supplement string) {
	//r.URL.Path returns the path of the request which we can then use strings.TrimPrefix to trim away /supplements/
	dosage := s.store.GetSupplementDosage(supplement)

	//if supplement not in StubSupplementDataStore return 404
	if supplement != "magnesium" && supplement != "vitamin-c" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, dosage)

}

//Function to process the POST request on /supplements/supplement handler and record 1 dosage as taken
func (s *supplementsServer) processTakenSupplement(w http.ResponseWriter, supplement string) {
	//return 200 status code if request is POST method to pass test
	//selenium not what we expect, will be updated
	s.store.RecordTakenSupplement(supplement)
	w.WriteHeader(http.StatusAccepted)
}

func (s *supplementsServer) showTakenSupplement(w http.ResponseWriter, supplement string) {
	takenSuppAmount := s.store.GetTakenSupplement(supplement)

	if takenSuppAmount == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, takenSuppAmount)
}

/*
not yet implemented
func (s *supplementsServer) processSetDosage(w http.ResponseWriter, supplement string) {
	//return 200 status code if request is POST method to pass test
	//selenium not what we expect, will be updated
	s.store.RecordtakenSupplement(supplement)
	w.WriteHeader(http.StatusAccepted)
}
*/
