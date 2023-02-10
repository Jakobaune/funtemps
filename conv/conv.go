package conv

var Celsius float64
var Fahrenheit float64
var Kelvin float64

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    fahrenheitToCelsius
    CelsiusToFahrenheit
    KelvinTofahrenheit
    ...
*/

// Konverterer fahrenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	Celsius = (Fahrenheit - 32) * (-1.8)

	return Celsius
}

func CelsiusToFahrenheit(value float64) float64 {
	Fahrenheit := value*(1.8) + 32

	return Fahrenheit
}

func CelsiusToKelvin(value float64) float64 {
	Kelvin = value + 273.15

	return Kelvin
}
func KelvinToCelsius(value float64) float64 {
	Celsius = value - 273.15

	return Celsius
}

func KelvinToFahrenheit(value float64) float64 {
	Fahrenheit := (value-273.15)*1.8 + 32

	return Fahrenheit
}
func FahrenheitToKelvin(value float64) float64 {
	Kelvin := (value-32)*1.8 + 273.15

	return Kelvin
}
