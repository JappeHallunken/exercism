// Package weather provides a weather forecast
// for a given city.
package weather

// CurrentCondition is the current weather condition as a string.
var CurrentCondition string

// CurrentLocation is the current City location as a string.
var CurrentLocation string

// Forecast takes a city and the current weather condition and returns a string with both values combined. SO you have the current weather in the location provided.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
