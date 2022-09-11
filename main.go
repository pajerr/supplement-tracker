package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/supplements", supplementsHandler)
	http.ListenAndServe(":8088", nil)
}
