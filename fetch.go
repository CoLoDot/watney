package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fetchInSightWeather() string {
	apiKey := getEnvVariable("API_KEY")

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
