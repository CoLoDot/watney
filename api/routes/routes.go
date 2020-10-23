package routes

import (
	"api/fetch"
	"log"
	"net/http"
)

// GetMarsWeather writes JSON response of InsightWeather
func GetMarsWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fetch.InsightWeather()))
	log.Println("Watney -> Mars Weather")
}

// GetAPOD writes JSON response of Apod
func GetAPOD(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fetch.Apod()))
	log.Println("Watney -> APOD")
}

// NotFound writes JSON response in an unknown endpoint
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
	log.Println("Watney -> nothing here")
}
