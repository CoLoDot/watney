package routes

import (
	"api/fetch"
	"log"
	"net/http"
	"time"
)

// GetMarsWeather writes JSON response of InsightWeather
func GetMarsWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fetch.InsightWeather()))
	log.Println("Watney -> Mars Weather")
}

// GetAPOD writes JSON response of Apod
func GetAPOD(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	today := time.Now().Format("2006-01-02")
	w.Write([]byte(fetch.Apod(today)))
	log.Println("Watney -> APOD")
}

// GetNeo writes JSON response of Neo
func GetNeo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	today := time.Now().Format("2006-01-02")
	w.Write([]byte(fetch.Neo(today)))
	log.Println("Watney -> Neo")
}

// NotFound writes JSON response in an unknown endpoint
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
	log.Println("Watney -> nothing here")
}
