package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fromKelvinToCelius(t *testing.T) {
	type testCase struct {
		kelvin          float32
		expectedCelsius float32
	}

	testCases := []testCase{
		{
			kelvin:          0,
			expectedCelsius: -273.15,
		},
		{
			kelvin:          273.15,
			expectedCelsius: 0,
		},
	}

	for _, tcase := range testCases {
		t.Run("", func(t *testing.T) {
			actual := fromKelvinToCelius(tcase.kelvin)
			assert.Equal(t, tcase.expectedCelsius, actual)
		})
	}
}

func Test_fromKelvinToFahrenheit(t *testing.T) {
	type testCase struct {
		kelvin             float32
		expectedFahrenheit float32
	}

	testCases := []testCase{
		{
			kelvin:             284,
			expectedFahrenheit: 51.53001,
		},
		{
			kelvin:             273.15,
			expectedFahrenheit: 32,
		},
	}

	for _, tcase := range testCases {
		t.Run("", func(t *testing.T) {
			actual := fromKelvinToFahrenheit(tcase.kelvin)
			assert.Equal(t, tcase.expectedFahrenheit, actual)
		})
	}
}
