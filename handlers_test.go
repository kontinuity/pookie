package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLoginPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(loginPageHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), "Login") {
		t.Errorf("handler should contain 'Login' in response body")
	}
}

func TestLoginHandlerSuccess(t *testing.T) {
	form := url.Values{}
	form.Add("username", "admin")
	form.Add("password", "password")

	req, err := http.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(loginHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}

	if location := rr.Header().Get("Location"); location != "/dashboard" {
		t.Errorf("handler returned wrong redirect location: got %v want %v",
			location, "/dashboard")
	}
}

func TestLoginHandlerFailure(t *testing.T) {
	form := url.Values{}
	form.Add("username", "wrong")
	form.Add("password", "credentials")

	req, err := http.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(loginHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), "Invalid username or password") {
		t.Errorf("handler should contain error message in response body")
	}
}

func TestLoginHandlerGetRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(loginHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}

	if location := rr.Header().Get("Location"); location != "/" {
		t.Errorf("handler returned wrong redirect location: got %v want %v",
			location, "/")
	}
}

func TestDashboardHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/dashboard", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(dashboardHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), "Welcome to Dashboard") {
		t.Errorf("handler should contain 'Welcome to Dashboard' in response body")
	}
}