package service

type Temperature struct {
	Kelvin     float32 `json:"kelvin"`
	Celsius    float32 `json:"celsius"`
	Fahrenheit float32 `json:"fahrenheit"`
}

type CityTemperatureData struct {
	City        string      `json:"city"`
	Temperature Temperature `json:"temperature"`
}

func generateCityTemperatureData(cityName string, kelvinTemperature float32) CityTemperatureData {
	return CityTemperatureData{
		City: cityName,
		Temperature: Temperature{
			Kelvin:     kelvinTemperature,
			Celsius:    fromKelvinToCelius(kelvinTemperature),
			Fahrenheit: fromKelvinToFahrenheit(kelvinTemperature),
		},
	}
}

func fromKelvinToCelius(kelvinTemperature float32) float32 {
	const ZERO_CELSIUS_IN_KELVIN = float32(273.15)
	return kelvinTemperature - ZERO_CELSIUS_IN_KELVIN
}

func fromKelvinToFahrenheit(kelvinTemperature float32) float32 {
	celsiusTemperature := fromKelvinToCelius(kelvinTemperature)
	const ZERO_FAHRENHEIT_IN_CELSIUS = 32
	const CELSIUS_TO_FAHRENHEIT_SCALE_FACTOR = float32(9) / float32(5)
	return celsiusTemperature*CELSIUS_TO_FAHRENHEIT_SCALE_FACTOR + ZERO_FAHRENHEIT_IN_CELSIUS
}
