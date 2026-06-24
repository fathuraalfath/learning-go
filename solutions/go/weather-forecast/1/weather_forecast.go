// Package weather provide tools for weather.
package weather

var (
    // CurrentCondition represent a certain condition.
	CurrentCondition string
    // CurrentLocation represent a certain location.
	CurrentLocation  string
)

// Forecast returns a string value equal current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
