package main

import (
	"log"
	"net/http"
)

/*
type InMemorySupplementDataStore struct{}

//we have not passed in a PlayerStore, so we need to hardcode response for now
func (i *InMemorySupplementDataStore) GetSupplementDosage(name string) int {
	return 123
}

func (i *InMemorySupplementDataStore) RecordTakenSupplement(name string) {
}

func (i *InMemorySupplementDataStore) GetTakenSupplement(name string) int {
	return 9
}
*/

//main now uses type from in_memory_supplement_store.go
func main() {
	server := &supplementsServer{NewInMemorySupplementStore()}
	log.Fatal(http.ListenAndServe(":5050", server))
}
