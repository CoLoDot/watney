package fetch

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("ENV", "TEST")
	os.Setenv("API_KEY", "DEMO_KEY")
}

func TestApod(t *testing.T) {
	actual := Apod("2020-11-01")
	expected := `{ "title": "In the Center of the Trifid Nebula", "url": "https://apod.nasa.gov/apod/image/2011/Trifid_HubbleGendler_2400.jpg", "description": "What's happening at the center of the Trifid Nebula? Three prominent dust lanes that give the Trifid its name all come together. Mountains of opaque dust appear near the bottom, while other dark filaments of dust are visible threaded throughout the nebula.  A single massive star visible near the center causes much of the Trifid's glow.  The Trifid, cataloged as M20, is only about 300,000 years old, making it among the youngest emission nebulas known.  The star forming nebula lies about 9,000 light years away toward the constellation of the Archer (Sagittarius). The region pictured here spans about 10 light years.  The featured image is a composite with luminance taken from an image by the 8.2-m ground-based Subaru Telescope, detail provided by the 2.4-m orbiting Hubble Space Telescope, color data provided by Martin Pugh and image assembly and processing provided by Robert Gendler.   Portal Universe: Random APOD Generator"}`
	if actual != expected {
		t.Errorf("TestApod was incorrect, got: %s, want: %s.", actual, expected)
	}
}
