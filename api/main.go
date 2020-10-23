package main

import (
	"api/routes"
	"api/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Watney is running...")
	r := mux.NewRouter()

	marsweather := r.PathPrefix("/api/marsweather").Subrouter()
	marsweather.HandleFunc("", routes.GetMarsWeather).Methods(http.MethodGet)
	marsweather.HandleFunc("", routes.NotFound)

	apod := r.PathPrefix("/api/apod").Subrouter()
	apod.HandleFunc("", routes.GetAPOD).Methods(http.MethodGet)
	apod.HandleFunc("", routes.NotFound)

	log.Fatal(http.ListenAndServe(":"+utils.GetEnvVariable("PORT"), r))
}
