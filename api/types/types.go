package types

// Apod : Astronomy Picture Of the Day response struct required fields
type Apod struct {
	Explanation string `json:"explanation"`
	Title       string `json:"title"`
	Hdurl       string `json:"hdurl"`
}

// Neo : Asteroids near earth for today's date
type Neo struct {
	ElementCount int64 `json:"element_count"`
	Links        struct {
		Next string `json:"next"`
		Prev string `json:"prev"`
		Self string `json:"self"`
	} `json:"links"`
	NearEarthObjects struct {
		Today []struct {
			AbsoluteMagnitudeH float64 `json:"absolute_magnitude_h"`
			CloseApproachData  []struct {
				CloseApproachDate      string `json:"close_approach_date"`
				CloseApproachDateFull  string `json:"close_approach_date_full"`
				EpochDateCloseApproach int64  `json:"epoch_date_close_approach"`
				MissDistance           struct {
					Astronomical string `json:"astronomical"`
					Kilometers   string `json:"kilometers"`
					Lunar        string `json:"lunar"`
					Miles        string `json:"miles"`
				} `json:"miss_distance"`
				OrbitingBody     string `json:"orbiting_body"`
				RelativeVelocity struct {
					KilometersPerHour   string `json:"kilometers_per_hour"`
					KilometersPerSecond string `json:"kilometers_per_second"`
					MilesPerHour        string `json:"miles_per_hour"`
				} `json:"relative_velocity"`
			} `json:"close_approach_data"`
			EstimatedDiameter struct {
				Feet struct {
					EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
					EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
				} `json:"feet"`
				Kilometers struct {
					EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
					EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
				} `json:"kilometers"`
				Meters struct {
					EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
					EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
				} `json:"meters"`
				Miles struct {
					EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
					EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
				} `json:"miles"`
			} `json:"estimated_diameter"`
			ID                             string `json:"id"`
			IsPotentiallyHazardousAsteroid bool   `json:"is_potentially_hazardous_asteroid"`
			IsSentryObject                 bool   `json:"is_sentry_object"`
			Links                          struct {
				Self string `json:"self"`
			} `json:"links"`
			Name           string `json:"name"`
			NasaJplURL     string `json:"nasa_jpl_url"`
			NeoReferenceID string `json:"neo_reference_id"`
		} `json:"today"`
	} `json:"near_earth_objects"`
}

// NeoObject : Neo object type
type NeoObject map[string]interface{}
