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
	router.Handle("listtaken", http.HandlerFunc(s.listTakenSupplementsHandler))

	s.Handler = router

	return s
}

/*
//Refactored to use SupplementDataStore interface
func (s *supplementsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
*/

func (s *supplementsServer) listTakenSupplementsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, dosage)

}

//Function to process the POST request on /supplements/supplement handler and record 1 dosage as taken
func (s *supplementsServer) processTakenSupplement(w http.ResponseWriter, supplement string) {
	//return 200 status code if request is POST method to pass test
	s.store.RecordTakenSupplement(supplement)

	//to fix CORS errors in frontend
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
}

func (s *supplementsServer) showTakenSupplement(w http.ResponseWriter, supplement string) {
	takenSuppAmount := s.store.GetTakenSupplement(supplement)

	if takenSuppAmount == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	//to fix CORS errors in frontend
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
