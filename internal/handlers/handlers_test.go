package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSimpleHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SimpleHelloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Hello, NDC! Pipeline Test v2.1.0 - Refactored Structure"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response HelloResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("could not unmarshal response: %v", err)
	}

	expectedMessage := "Hello, NDC! Pipeline Test v2.1.0 - Refactored Structure"
	if response.Message != expectedMessage {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, expectedMessage)
	}

	if response.Version != "v2.1.0" {
		t.Errorf("handler returned unexpected version: got %v want %v",
			response.Version, "v2.1.0")
	}
}

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response HealthResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("could not unmarshal response: %v", err)
	}

	if response.Status != "healthy" {
		t.Errorf("handler returned unexpected status: got %v want %v",
			response.Status, "healthy")
	}

	if response.Version != "v2.1.0" {
		t.Errorf("handler returned unexpected version: got %v want %v",
			response.Version, "v2.1.0")
	}
}

func TestDashboard(t *testing.T) {
	req, err := http.NewRequest("GET", "/dashboard", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Dashboard)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "text/html" {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, "text/html")
	}

	body := rr.Body.String()
	if !strings.Contains(body, "Jenkins CI/CD Demo") {
		t.Errorf("handler body should contain 'Jenkins CI/CD Demo'")
	}

	if !strings.Contains(body, "Pipeline Stages Completed") {
		t.Errorf("handler body should contain 'Pipeline Stages Completed'")
	}
}

func TestMetrics(t *testing.T) {
	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Metrics)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, "application/json")
	}

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("could not unmarshal response: %v", err)
	}

	if response["go_version"] != "1.21" {
		t.Errorf("handler returned unexpected go_version: got %v want %v",
			response["go_version"], "1.21")
	}

	if response["status"] != "running" {
		t.Errorf("handler returned unexpected status: got %v want %v",
			response["status"], "running")
	}

	stages, ok := response["pipeline_stages"].([]interface{})
	if !ok {
		t.Errorf("pipeline_stages should be an array")
	}

	if len(stages) != 6 {
		t.Errorf("pipeline_stages should have 6 stages, got %d", len(stages))
	}
}