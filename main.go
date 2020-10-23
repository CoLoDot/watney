package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	marsweather := r.PathPrefix("/api/marsweather").Subrouter()
	marsweather.HandleFunc("", getMarsWeather).Methods(http.MethodGet)
	marsweather.HandleFunc("", notFound)

	apod := r.PathPrefix("/api/apod").Subrouter()
	apod.HandleFunc("", getAPOD).Methods(http.MethodGet)
	apod.HandleFunc("", notFound)

	log.Fatal(http.ListenAndServe(":"+getEnvVariable("PORT"), r))
}
