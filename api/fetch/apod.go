package fetch

import (
	"api/types"
	"encoding/json"
	"fmt"

	"api/utils"
)

// Apod fetches Astronomy Picture of Day NASA Api data
func Apod() string {
	apiKey := utils.GetEnvVariable("API_KEY")
	query := "https://api.nasa.gov/planetary/apod?api_key=" + apiKey
	response := utils.GetResponse(query)

	var apodData types.Apod
	errBody := json.Unmarshal(response, &apodData)
	if errBody != nil {
		panic(errBody)
	}

	return fmt.Sprintf(`{ "title": "%s", "url": "%s", "description": "%s"}`, apodData.Title, apodData.Hdurl, apodData.Explanation)
}
