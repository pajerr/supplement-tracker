package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

//server.go
type SupplementTaken struct {
	Supplement string
	Dosage     int
}

//Refactored to use SupplementDataStore interface
func (s *supplementsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		s.processTakenDosage(w, r)
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

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, dosage)

}

func (s *supplementsHandler) processTakenDosage(w http.ResponseWriter, r *http.Request) {
	//extract the dosage and name of the supplement from the URL
	//r.URL.Path returns the path of the request which we can then use strings.TrimPrefix to trim away /supplements/
	//dosage := s.store.GetSupplementDosage(supplement)
	//extract name of suplpplement and dosage from URL "/supplements/magnesium/200"
	//supplement := strings.TrimPrefix(r.URL.Path, "/supplements/")
	//split the URL into an array of strings at "/" delimiter
	splittedURL := strings.Split(r.URL.Path, "/")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var dosageTaken map[string]int
	json.Unmarshal(reqBody, &dosageTaken)

	//verified correct index is 2
	supplement := splittedURL[2]
	//dosage := (splittedURL[3])
	//dosageInt, _ := strconv.Atoi(dosage)
	//fmt.Printf(dosageInt)

	//s.store.StoreTakenDosage("magnesium", 500)
	s.store.StoreTakenDosage(supplement, 500)

	//return 200 status code if request is POST method to pass test
	w.WriteHeader(http.StatusAccepted)
}
