package main

// Apod : Astronomy Picture Of the Day response struct required fields
type Apod struct {
	Explanation string `json:"explanation"`
	Title       string `json:"title"`
	Hdurl       string `json:"hdurl"`
}
