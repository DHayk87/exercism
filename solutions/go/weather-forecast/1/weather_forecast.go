// Package weather provides functionality to get the current weather for a location.
package weather

// CurrentCondition is a string that represents the current weather condition.
var CurrentCondition string

// CurrentLocation is a string that represents the current location.
var CurrentLocation string

// Forecast returns the current weather for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}