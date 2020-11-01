package fetch

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("ENV", "TEST")
	os.Setenv("API_KEY", "DEMO_KEY")
}

func TestNeo(t *testing.T) {
	actual := Neo("2020-11-01")
	expected := `{ "totalObjects": 12, "date": "2020-11-01", "nearEarthObjects": [
{
"closeApproachDate": "2020-Nov-01 14:08",
"estimatedDiameterInMeters": 365.97,
"id": "54054500",
"name": "(2020 SK1)",
"nasaJplUrl": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=54054500",
"relativeVelocityKMPerSecond": 4.7
},
{
"closeApproachDate": "2020-Nov-01 06:55",
"estimatedDiameterInMeters": 35.11,
"id": "54016669",
"name": "(2020 HO3)",
"nasaJplUrl": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=54016669",
"relativeVelocityKMPerSecond": 8.03
},
{
"closeApproachDate": "2020-Nov-01 08:35",
"estimatedDiameterInMeters": 510.32,
"id": "54053756",
"name": "(2020 QA7)",
"nasaJplUrl": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=54053756",
"relativeVelocityKMPerSecond": 3.94
},
{
"closeApproachDate": "",
"estimatedDiameterInMeters": 392.68,
"id": "3257066",
"name": "(2004 TL19)",
"nasaJplUrl": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3257066",
"relativeVelocityKMPerSecond": 8.1
},
{
"closeApproachDate": "2020-Nov-01 14:24",
"estimatedDiameterInMeters": 182.32,
"id": "54079790",
"name": "(2020 UN6)",
"nasaJplUrl": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=54079790",
"relativeVelocityKMPerSecond": 13.25
},
{
"closeApproachDate": "2020-Nov-01 06:35",
"estimatedDiameterInMeters": 85.91,
"id": "3788123",
"name": "(2017 UH8)",
"nasaJplUrl": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3788123",
"relativeVelocityKMPerSecond": 19.77
},
{
"closeApproachDate": "2020-Nov-01 01:15",
"estimatedDiameterInMeters": 392.68,
"id": "2474613",
"name": "474613 (2004 TL19)",
"nasaJplUrl": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=2474613",
"relativeVelocityKMPerSecond": 8.1
},
{
"closeApproachDate": "2020-Nov-01 09:24",
"estimatedDiameterInMeters": 156.33,
"id": "3872623",
"name": "(2019 SK7)",
"nasaJplUrl": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3872623",
"relativeVelocityKMPerSecond": 13.5
},
{
"closeApproachDate": "2020-Nov-01 20:39",
"estimatedDiameterInMeters": 73.42,
"id": "54016363",
"name": "(2020 FY4)",
"nasaJplUrl": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=54016363",
"relativeVelocityKMPerSecond": 6.88
},
{
"closeApproachDate": "2020-Nov-01 18:20",
"estimatedDiameterInMeters": 122.02,
"id": "54054685",
"name": "(2020 ST5)",
"nasaJplUrl": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=54054685",
"relativeVelocityKMPerSecond": 2.78
},
{
"closeApproachDate": "2020-Nov-01 05:57",
"estimatedDiameterInMeters": 65.56,
"id": "54076362",
"name": "(2020 UF5)",
"nasaJplUrl": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=54076362",
"relativeVelocityKMPerSecond": 2.82
},
{
"closeApproachDate": "2020-Nov-01 01:34",
"estimatedDiameterInMeters": 192.5,
"id": "54076933",
"name": "(2020 UV5)",
"nasaJplUrl": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=54076933",
"relativeVelocityKMPerSecond": 10.61
}
]}`
	if actual != expected {
		t.Errorf("TestNeo was incorrect, got: %s, want: %s.", actual, expected)
	}
}
