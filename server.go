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
type supplementsHandler struct {
	store SupplementDataStore
}

//Refactored to use SupplementDataStore interface
func (s *supplementsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//r.URL.Path returns the path of the request which we can then use strings.TrimPrefix to trim away /supplements/
	supplement := strings.TrimPrefix(r.URL.Path, "/supplements/")

	switch r.Method {
	//if client sends GET request
	case http.MethodGet:
		//if supplement is not "vitamin-c" or "magensium"
		//if supplement != "vitamin-c" && supplement != "magnesium" {
		//	w.WriteHeader(http.StatusNotFound)
		//}

		if supplement == "vitamin-c" {
			w.WriteHeader(http.StatusOK)
			//write dosage of Vitamin C to the response
			fmt.Fprintf(w, "%v", s.store.GetSupplementDosage("vitamin-c"))
		} else if supplement == "magnesium" {
			w.WriteHeader(http.StatusOK)
			//write dosage of Magnesium to the response
			fmt.Fprintf(w, "%v", s.store.GetSupplementDosage("magnesium"))
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Supplement not found")
		}
	}
}
