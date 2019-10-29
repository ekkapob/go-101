package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	if greet("Joe") != "Hello Joe" {
		t.Errorf("expected Hello Joe but got %v", greet("Joe"))
	}
}

func TestHomeHandler(t *testing.T) {
	t.Run("home handler", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/", nil)
		if err != nil {
			t.Error(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(homeHandler)
		handler.ServeHTTP(resp, req)

		if status := resp.Code; status != http.StatusOK {
			t.Errorf("expected %v but got %v", http.StatusOK, status)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		expectedHTML := "Type your name"
		if match, _ := regexp.Match(expectedHTML, body); !match {
			t.Errorf("expected %v but got %v", expectedHTML, string(body))
		}
	})

	t.Run("not found url", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/not-found", nil)
		if err != nil {
			t.Error(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(homeHandler)
		handler.ServeHTTP(resp, req)

		if status := resp.Code; status != http.StatusFound {
			t.Errorf("expected %v but got %v", http.StatusFound, status)
		}

	})
}

func TestGreetHandler(t *testing.T) {
	t.Run("greet handler", func(t *testing.T) {
		data := url.Values{"name": []string{"Jordan"}}

		req, err := http.NewRequest(http.MethodPost, "/greet", strings.NewReader(data.Encode()))
		if err != nil {
			t.Error(err)
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(greetHandler)
		handler.ServeHTTP(resp, req)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		if status := resp.Code; status != http.StatusOK {
			t.Errorf("expected %v but got %v", http.StatusOK, status)
		}

		expectedHTML := "Hello Jordan"
		if match, _ := regexp.Match(expectedHTML, body); !match {
			t.Errorf("expected %v but got %v", expectedHTML, string(body))
		}
	})
}
