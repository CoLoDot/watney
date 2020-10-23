package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

// GetResponse helper for RESTApi
func GetResponse(query string) []byte {
	resp, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
