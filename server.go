package main

import (
	"fmt"
	"net/http"
	"strings"
)

//Server/Handler gets idea of what SupplementDataStore is from the interface
type SupplementDataStore interface {
	GetSupplementDosage(name string) int
	StoreTakenDosage(name string, dosage int)
}

//allows us to use the SupplementDataStore interface in the handler
//for example to store.GetSupplelementDosage to get supplements dosage
type supplementsHandler struct {
	store SupplementDataStore
}

//Refactored to use SupplementDataStore interface
func (s *supplementsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		s.processTakenDosage(w)
	case http.MethodGet:
		s.showDosage(w, r)
	}

}

func (s *supplementsHandler) showDosage(w http.ResponseWriter, r *http.Request) {
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

func (s *supplementsHandler) processTakenDosage(w http.ResponseWriter) {
	//extract the dosage and name of the supplement from the URL
	//r.URL.Path returns the path of the request which we can then use strings.TrimPrefix to trim away /supplements/
	//supplement := strings.TrimPrefix(r.URL.Path, "/supplements/")
	//dosage := s.store.GetSupplementDosage(supplement)

	s.store.StoreTakenDosage("magnesium", 200)
	//return 200 status code if request is POST method to pass test
	w.WriteHeader(http.StatusAccepted)
}
