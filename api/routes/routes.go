package routes

import (
	"api/fetch"
	"net/http"
)

// GetMarsWeather writes JSON response of FetchInSightWeather
func GetMarsWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fetch.InsightWeather()))
}

// GetAPOD writes JSON response of FetchAPOD
func GetAPOD(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fetch.Apod()))
}

// NotFound write JSON response in an unknown endpoint
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
