package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func fetchInSightWeather() string {
	apiKey := os.Getenv("API_KEY")

	resp, err := http.Get("https://api.nasa.gov/insight_weather/?api_key=" + apiKey + "&feedtype=json&ver=1.0")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	errBody := json.Unmarshal(body, &data)
	if errBody != nil {
		panic(errBody)
	}

	solKeysList := data["sol_keys"].([]interface{})

	solkey := solKeysList[len(solKeysList)-1].(string)
	solData := data[solkey].(map[string]interface{})

	solDataAT := solData["AT"].(map[string]interface{})
	solDataATav := solDataAT["av"].(float64)
	solDataATmx := solDataAT["mx"].(float64)
	solDataATmn := solDataAT["mn"].(float64)

	solDataFirstUTC := solData["First_UTC"].(string)
	solDataLastUTC := solData["Last_UTC"].(string)
	solDataSeason := solData["Season"].(string)

	return fmt.Sprintf(`{"sol": "%s", "avTemp": %f, "maxTemp": %f, "minTemp": %f, "firstUTC": "%s", "lastUTC": "%s", "season": "%s"}`, solkey, solDataATav, solDataATmx, solDataATmn, solDataFirstUTC, solDataLastUTC, solDataSeason)
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fetchInSightWeather()))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", notFound)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
