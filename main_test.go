package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVitaminCDosage(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/vitaminc", nil)
	response := httptest.NewRecorder()

	supplementsHandler(response, request)

	t.Run("Return Vitamin C dosage", func(t *testing.T) {
		got := response.Body.String()
		want := "500"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
