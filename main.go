package main

import (
	"flag"
	"fmt"
	"funtemps/conv"
	"strconv"
)

// Definerer flag-variablene i hoved-"scope"
var F float64
var K float64
var C float64
var out string
var funfacts string

func init() {

	flag.Float64Var(&F, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&C, "C", 0.0, "Temperatur i Celsius")
	flag.Float64Var(&K, "K", 0.0, "Temperatur i Kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")

}

func main() {
	flag.Parse()

	// Parse input temperature and its scale
	inputScale := flag.Arg(0)
	inputValue, err := strconv.ParseFloat(flag.Arg(1), 64)
	if err != nil {
		fmt.Println("Invalid input temperature value")
		return
	}

	// Parse output scale
	outputScale := flag.Lookup("out").Value.String()

	// Convert temperature
	var outputValue float64
	switch {
	case inputScale == "C" && outputScale == "F":
		outputValue = conv.CelsiusToFahrenheit(inputValue)
	case inputScale == "F" && outputScale == "C":
		outputValue = conv.FahrenheitToCelsius(inputValue)
	case inputScale == "C" && outputScale == "K":
		outputValue = conv.CelsiusToKelvin(inputValue)
	case inputScale == "K" && outputScale == "C":
		outputValue = conv.KelvinToCelsius(inputValue)
	case inputScale == "F" && outputScale == "K":
		outputValue = conv.FahrenheitToKelvin(inputValue)
	case inputScale == "K" && outputScale == "F":
		outputValue = conv.KelvinToFahrenheit(inputValue)
	default:
		fmt.Println("Invalid input/output scale combination")
		return
	}

	// Print result
	fmt.Printf("%.2f°%s is %.2f°%s\n", inputValue, inputScale, outputValue, outputScale)
}
