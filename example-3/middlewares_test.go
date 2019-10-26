package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddHeader(t *testing.T) {
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "", nil)
	h := func(w http.ResponseWriter, r *http.Request) {
		expectedHeaderValue := "foobar"
		if r.Header.Get("custom-header") != expectedHeaderValue {
			t.Errorf("expected header of %v but got %v",
				expectedHeaderValue,
				r.Header.Get("custom-header"),
			)
		}
	}
	AddHeader(h)(resp, req)
}
