package fetch

import (
	"api/utils"
	"encoding/json"
	"fmt"
)

// InsightWeather fetches Mars InsightWeather NASA Api data
func InsightWeather() string {
	apiKey := utils.GetEnvVariable("API_KEY")
	query := "https://api.nasa.gov/insight_weather/?api_key=" + apiKey + "&feedtype=json&ver=1.0"
	response := utils.GetResponse(query)

	var insightWeatherData map[string]interface{}
	errBody := json.Unmarshal(response, &insightWeatherData)
	if errBody != nil {
		panic(errBody)
	}

	solKeysList := insightWeatherData["sol_keys"].([]interface{})

	solkey := solKeysList[len(solKeysList)-1].(string)
	solData := insightWeatherData[solkey].(map[string]interface{})

	solDataAT := solData["AT"].(map[string]interface{})
	solDataATav := solDataAT["av"].(float64)
	solDataATmx := solDataAT["mx"].(float64)
	solDataATmn := solDataAT["mn"].(float64)

	solDataFirstUTC := solData["First_UTC"].(string)
	solDataLastUTC := solData["Last_UTC"].(string)
	solDataSeason := solData["Season"].(string)

	return fmt.Sprintf(`{"sol": "%s", "avTemp": %f, "maxTemp": %f, "minTemp": %f, "firstUTC": "%s", "lastUTC": "%s", "season": "%s"}`, solkey, solDataATav, solDataATmx, solDataATmn, solDataFirstUTC, solDataLastUTC, solDataSeason)
}
