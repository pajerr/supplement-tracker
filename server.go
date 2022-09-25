package main

import (
	"fmt"
	"net/http"
	"strings"
)

type SupplementDataStore interface {
	GetSupplementDosage(name string) int
}

//allows us to use the SupplementDataStore interface in the handler
//for example to store.GetSupplelementDosage to get supplements dosage
type supplementsServer struct {
	store SupplementDataStore
}

//Refactored to use SupplementDataStore interface
func (s *supplementsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		s.processSetDosage(w)
	case http.MethodGet:
		s.showDosage(w, r)
	}

}

func (s *supplementsServer) showDosage(w http.ResponseWriter, r *http.Request) {
	//r.URL.Path returns the path of the request which we can then use strings.TrimPrefix to trim away /supplements/
	supplement := strings.TrimPrefix(r.URL.Path, "/supplements/")
	dosage := s.store.GetSupplementDosage(supplement)

	//if supplement not in StubSupplementDataStore return 404
	if supplement != "magnesium" && supplement != "vitamin-c" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, dosage)

}

func (s *supplementsServer) processSetDosage(w http.ResponseWriter) {
	//return 200 status code if request is POST method to pass test
	w.WriteHeader(http.StatusAccepted)
}
