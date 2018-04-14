// Package space exports several space-related timing functions
package space

// Planet is a lil string baby
type Planet string

// Orbital years in seconds (by planet)
const (
	Earth   float64 = 31557600
	Mercury float64 = Earth * 0.2408467
	Venus   float64 = Earth * 0.61519726
	Mars    float64 = Earth * 1.8808158
	Jupiter float64 = Earth * 11.862615
	Saturn  float64 = Earth * 29.447498
	Uranus  float64 = Earth * 84.016846
	Neptune float64 = Earth * 164.79132
)

// Age calculates your age on a given planet based on orbital years.
func Age(seconds float64, planet Planet) float64 {
	var age float64

	switch planet {
	case "Earth":
		age = seconds / Earth
	case "Mercury":
		age = seconds / Mercury
	case "Venus":
		age = seconds / Venus
	case "Mars":
		age = seconds / Mars
	case "Jupiter":
		age = seconds / Jupiter
	case "Saturn":
		age = seconds / Saturn
	case "Uranus":
		age = seconds / Uranus
	case "Neptune":
		age = seconds / Neptune
	}

	return age
}
