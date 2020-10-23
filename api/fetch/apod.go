package fetch

import (
	"api/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"api/utils"
)

// Apod fetches Astronomy Picture of Day NASA Api data
func Apod() string {
	apiKey := utils.GetEnvVariable("API_KEY")

	resp, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=" + apiKey)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var apodData types.Apod
	errBody := json.Unmarshal(body, &apodData)
	if errBody != nil {
		panic(errBody)
	}

	return fmt.Sprintf(`{ "title": "%s", "url": "%s", "description": "%s"}`, apodData.Title, apodData.Hdurl, apodData.Explanation)
}
