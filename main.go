package main

import (
	"log"
	"net/http"
)

func main() {
	server := &supplementsHandler{}
	log.Fatal(http.ListenAndServe(":5050", server))
}
