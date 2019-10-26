package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

type CheckHeaderFunc func(resp *http.Request) bool

func TestHandlers(t *testing.T) {

	tests := []struct {
		route              string
		expectedMsg        string
		expectedStatus     int
		redirectedLocation string
		checkHeader        CheckHeaderFunc
		handlerFunc        http.HandlerFunc
	}{
		{
			route:          "/",
			expectedMsg:    "This is Home",
			expectedStatus: http.StatusOK,
			handlerFunc:    RootHandler,
		},
		{
			route:          "/home",
			expectedMsg:    "Hello World",
			expectedStatus: http.StatusOK,
			handlerFunc:    HomeHandler,
		},
		{
			route:          "/log",
			expectedMsg:    "Hello World with Middlewares",
			expectedStatus: http.StatusOK,
			handlerFunc:    Logger(LogHandler),
		},
		{
			route:          "/header",
			expectedStatus: http.StatusOK,
			checkHeader:    checkHeader("custom-header", "foobar"),
			handlerFunc:    AddHeader(Logger(LogHandler)),
		},
		{
			route:          "/authen/pass",
			expectedMsg:    "Congrats!",
			expectedStatus: http.StatusOK,
			handlerFunc:    AddHeader(Authen(Logger(PassAuthenHandler))),
		},
		{
			route:              "/authen/fail",
			expectedStatus:     http.StatusFound,
			redirectedLocation: "/",
			handlerFunc:        Authen(AddHeader(Logger(PassAuthenHandler))),
		},
		{
			route:          "/not-found-url",
			expectedMsg:    "Not Found",
			expectedStatus: http.StatusNotFound,
			handlerFunc:    RootHandler,
		},
	}

	for _, test := range tests {

		testName := fmt.Sprintf("verify successful result on route %v", test.route)
		t.Run(testName, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, test.route, nil)
			if err != nil {
				t.Error(err)
			}

			resp := httptest.NewRecorder()
			handler := http.HandlerFunc(test.handlerFunc)
			handler.ServeHTTP(resp, req)

			if status := resp.Code; status != test.expectedStatus {
				t.Errorf("expected %v but got %v", test.expectedStatus, status)
			}

			if test.checkHeader != nil && !test.checkHeader(req) {
				t.Errorf("expected header not match")
			}

			if test.redirectedLocation != "" {
				location, err := resp.Result().Location()
				if err != nil {
					t.Error(err)
				}
				if location.Path != test.redirectedLocation {
					t.Errorf("expected %v but got %v", test.redirectedLocation, location.Path)
				}
			}

			if test.expectedMsg != "" {
				matched, _ := regexp.Match(fmt.Sprint(test.expectedMsg), resp.Body.Bytes())
				if !matched {
					t.Errorf("expected %v but got %v", test.expectedMsg, resp.Body.String())
				}
			}
		})
	}

}

func BenchmarkAuthen(b *testing.B) {
	b.Run("benchmark authen pass", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req, err := http.NewRequest(http.MethodGet, "/authen/pass", nil)
			if err != nil {
				b.Error(err)
			}
			resp := httptest.NewRecorder()
			handler := http.HandlerFunc(AddHeader(Authen(Logger(PassAuthenHandler))))
			handler.ServeHTTP(resp, req)
			matched, _ := regexp.Match(fmt.Sprint("Congrats!"), resp.Body.Bytes())
			if !matched {
				b.Errorf("expected %v but got %v", "Congrats!", resp.Body.String())
			}
		}
	})
}

func checkHeader(key, value string) CheckHeaderFunc {
	return func(req *http.Request) bool {
		if req.Header.Get(key) == value {
			return true
		}
		return false
	}
}
