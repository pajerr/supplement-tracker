package main

import (
	"fmt"
	"net/http"
)

//var supplementList []Supplement

type Supplement struct {
	Name   string `json:"Name"`
	Dosage int    `json:"Dosage"`
	Unit   string `json:"Unit"`
}

func supplementsHandler(w http.ResponseWriter, r *http.Request) {
	//hardcoded test data
	testVitaminC := Supplement{Name: "Vitamin C", Dosage: 500, Unit: "mg"}
	switch r.Method {
	//if client sends GET request
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		//write dosage of Vitamin C to the response
		fmt.Fprintf(w, "%v", testVitaminC.Dosage)
	}
}

func main() {
	http.HandleFunc("/supplements", supplementsHandler)
	http.ListenAndServe(":8088", nil)
}
