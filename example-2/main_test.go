package main

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestHome(t *testing.T) {
	t.Run("it should return Hello World message", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Error(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(HomeHandler)
		handler.ServeHTTP(resp, req)
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
		matched, _ := regexp.Match(`(?i)hello world`, resp.Body.Bytes())
		if !matched {
			t.Errorf("expected 'hello world' but got %v", resp.Body.String())
		}

	})
}
