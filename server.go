package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Supplement struct {
	Name   string `json:"Name"`
	Dosage int    `json:"Dosage"`
	Unit   string `json:"Unit"`
}

func supplementsHandler(w http.ResponseWriter, r *http.Request) {
	//hardcoded test data
	testVitaminC := Supplement{Name: "Vitamin C", Dosage: 500, Unit: "mg"}
	testMagnesium := Supplement{Name: "Magnesium", Dosage: 400, Unit: "mg"}

	//r.URL.Path returns the path of the request which we can then use strings.TrimPrefix to trim away /supplements/
	supplement := strings.TrimPrefix(r.URL.Path, "/supplements/")

	switch r.Method {
	//if client sends GET request
	case http.MethodGet:
		if supplement == "vitamin-c" {
			w.WriteHeader(http.StatusOK)
			//write dosage of Vitamin C to the response
			fmt.Fprintf(w, "%v", testVitaminC.Dosage)
		} else if supplement == "magnesium" {
			w.WriteHeader(http.StatusOK)
			//write dosage of Magnesium to the response
			fmt.Fprintf(w, "%v", testMagnesium.Dosage)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Supplement not found")
		}
	}
}
