package conv

import (
	"testing"
)

func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input  float64
		want   float64
		margin float64
	}

	tests := []test{
		{input: 134, want: 56.67, margin: 1},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if got < tc.want-tc.margin || got > tc.want+tc.margin {
			t.Errorf("input: %v, expected: %v +/- %v, but got: %v", tc.input, tc.want, tc.margin, got)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input  float64
		want   float64
		margin float64
	}

	tests := []test{
		{input: 20, want: 68, margin: 1},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if got < tc.want-tc.margin || got > tc.want+tc.margin {
			t.Errorf("input: %v, expected: %v +/- %v, but got: %v", tc.input, tc.want, tc.margin, got)
		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input  float64
		want   float64
		margin float64
	}

	tests := []test{
		{input: 50, want: 283.15, margin: 1},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if got < tc.want-tc.margin || got > tc.want+tc.margin {
			t.Errorf("input: %v, expected: %v +/- %v, but got: %v", tc.input, tc.want, tc.margin, got)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input  float64
		want   float64
		margin float64
	}

	tests := []test{
		{input: 283.15, want: 50, margin: 1},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if got < tc.want-tc.margin || got > tc.want+tc.margin {
			t.Errorf("input: %v, expected: %v +/- %v, but got: %v", tc.input, tc.want, tc.margin, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input  float64
		want   float64
		margin float64
	}

	tests := []test{
		{input: 1, want: -272.15, margin: 1},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if got < tc.want-tc.margin || got > tc.want+tc.margin {
			t.Errorf("input: %v, expected: %v +/- %v, but got: %v", tc.input, tc.want, tc.margin, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input  float64
		want   float64
		margin float64
	}

	tests := []test{
		{input: 20, want: 293.15, margin: 1},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if got < tc.want-tc.margin || got > tc.want+tc.margin {
			t.Errorf("input: %v, expected: %v +/- %v, but got: %v", tc.input, tc.want, tc.margin, got)
		}
	}
}
