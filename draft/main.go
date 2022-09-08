package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//type Supplement interface {
//Details() string
//}

var supplements []Supplement

type Supplement struct {
	Name   string `json:"Name"`
	Dosage int    `json:"Dosage"`
	Unit   string `json:"Unit"`
}

//function to add supplement to slice, will be replaced with DB
func addSupplement(supplement Supplement) {
	supplements = append(supplements, supplement)
}

func supplementsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//if client sends GET request
	case http.MethodGet:
		supplementsJSON, err := json.Marshal(supplements)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		//write the JSON to the response to the client
		w.Write(supplementsJSON)
	}
}

func main() {
	//Create empty map of Supplements and append supplements to it
	//will be replaced by DB
	//var supplements []Supplement
	supplements = append(supplements, Supplement{"Vitamin C", 500, "mg"})
	supplements = append(supplements, Supplement{"Vitamin D", 1000, "IU"})

	for range supplements {
		fmt.Println(supplements)
	}

	http.HandleFunc("/supplements", supplementsHandler)
	http.ListenAndServe(":8088", nil)
	//test to implement GET method
	//curl localhost:8088/supplements
	//[{"Name":"Vitamin C","Dosage":500,"Unit":"mg"},{"Name":"Vitamin D","Dosage":1000,"Unit":"IU"}]
}
