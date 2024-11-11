package space

type Planet string

func GetEarthYearsFromPlanet(planet Planet) float64 {
	switch planet {
	case "Mercury":
		return 0.2408467
	case "Venus":
		return 0.61519726
	case "Earth":
		return 1
	case "Mars":
		return 1.8808158
	case "Jupiter":
		return 11.862615
	case "Saturn":
		return 29.447498
	case "Uranus":
		return 84.016846
	case "Neptune":
		return 164.79132
	default:
		return -1
	}
}

func Age(seconds float64, planet Planet) float64 {
	earthYears := GetEarthYearsFromPlanet(planet)
	if earthYears < 0 {
		return earthYears
	}
	return (seconds / 31557600) / earthYears
}
