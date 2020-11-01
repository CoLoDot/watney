package fetch

import (
	"api/types"
	"api/utils"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

// Neo fetches objects near earth for current day
func Neo(today string) string {
	apiKey := utils.GetEnvVariable("API_KEY")
	query := "https://api.nasa.gov/neo/rest/v1/feed?start_date=" + today + "&end_date=" + today + "&api_key=" + apiKey
	response := utils.GetResponse(query)

	var neoDataCleaner map[string]interface{}
	errBodyCleaner := json.Unmarshal(response, &neoDataCleaner)
	if errBodyCleaner != nil {
		panic(errBodyCleaner)
	}
	neoDataCleaner["near_earth_objects"].(map[string]interface{})["today"] = neoDataCleaner["near_earth_objects"].(map[string]interface{})[today]
	r, _ := json.Marshal(neoDataCleaner)

	var neoData types.Neo
	errBody := json.Unmarshal(r, &neoData)
	if errBody != nil {
		panic(errBody)
	}

	nearEarthObjects := make([]types.NeoObject, 0)
	for _, val := range neoData.NearEarthObjects.Today {

		var relativeVelocityKMPerSecond float64
		if value, err := strconv.ParseFloat(val.CloseApproachData[0].RelativeVelocity.KilometersPerSecond, 64); err == nil {
			relativeVelocityKMPerSecond = value
		}

		nearEarthObjects = append(nearEarthObjects, types.NeoObject{
			"id":                          val.ID,
			"name":                        val.Name,
			"closeApproachDate":           val.CloseApproachData[0].CloseApproachDateFull,
			"estimatedDiameterInMeters":   math.Round(val.EstimatedDiameter.Meters.EstimatedDiameterMax*100) / 100,
			"relativeVelocityKMPerSecond": math.Round(relativeVelocityKMPerSecond*100) / 100,
			"nasaJplUrl":                  val.NasaJplURL,
		})
	}

	encodedNearEarthObjects, err := json.MarshalIndent(nearEarthObjects, "", "")
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(`{ "totalObjects": %d, "date": "%s", "nearEarthObjects": %s}`, neoData.ElementCount, today, encodedNearEarthObjects)
}
