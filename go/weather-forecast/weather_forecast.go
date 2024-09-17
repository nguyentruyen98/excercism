// Package weather provides a simple weather forecast package.
package weather

// CurrentCondition is the current weather condition.
var CurrentCondition string

// CurrentLocation is the location for which the weather forecast is being requested.
var CurrentLocation string

// Forecast returns the weather forecast for the specified city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
