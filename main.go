package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(supplementsHandler)
	log.Fatal(http.ListenAndServe(":5050", handler))
}
