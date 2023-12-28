package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	_ "gorm.io/driver/postgres"
)

func TestLoginSuccess(t *testing.T) {
	// Initialize
	router, _ := InitializeTestApp()

	// Request
	requestBody := strings.NewReader(`{"username": "testuser", "password": "password"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/login", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Read response
	response := recorder.Result()
	// Read response body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// Assertion
	// Status code is 200
	assert.Equal(t, response.StatusCode, http.StatusOK)
	assert.Equal(t, responseBody["status"], float64(200))
	// Response body contains access_token and value is not null
	assert.NotEqual(t, responseBody["data"].(map[string]interface{})["access_token"], nil)
}

func TestLoginFailed(t *testing.T) {
	// Initialize
	router, _ := InitializeTestApp()

	requestBody := strings.NewReader(`{"username": "testwrongpass", "password": "password"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/login", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Read response
	response := recorder.Result()
	// Read response body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// Assertion
	// Status code is 401
	assert.Equal(t, response.StatusCode, http.StatusUnauthorized)
	assert.Equal(t, responseBody["status"], float64(401))
	// Response body contains data and value is null
	assert.Equal(t, responseBody["data"], nil)
}
