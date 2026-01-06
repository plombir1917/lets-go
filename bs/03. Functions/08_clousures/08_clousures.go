package main

import "fmt"

func createTemperatureAdjuster() (func(change float64) float64, float64) {
	baseTemp := 90.0

	adjustTemperature := func(change float64) float64 {
		baseTemp += change
		return baseTemp
	}

	return adjustTemperature, baseTemp
}

func main() {
	adjustTemp, temp := createTemperatureAdjuster()
	fmt.Printf("Original temp %.1f\n", temp)

	adjustTemp(1.5)
	fmt.Printf("Adjusted temp %.1f\n", adjustTemp(1.5))

}
