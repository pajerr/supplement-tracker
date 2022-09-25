package main

import (
	"log"
	"net/http"
)

type InMemorySupplementDataStore struct{}

//we have not passed in a PlayerStore, so we need to hardcode response for now
func (i *InMemorySupplementDataStore) GetSupplementDosage(name string) int {
	return 123
}

//we have not passed in a PlayerStore, so we need to hardcode response for now
func (i *InMemorySupplementDataStore) RecordTakenDosage(name string) {
}

func main() {
	server := &supplementsServer{&InMemorySupplementDataStore{}}
	log.Fatal(http.ListenAndServe(":5050", server))
}
