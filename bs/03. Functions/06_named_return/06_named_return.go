package main

import "fmt"

func estimateBrewTime(cupsQty int, secondsPerCup int) (totalTime int) {
	totalTime = cupsQty * secondsPerCup
	return // naked return
}

func main() {
	brewTime := estimateBrewTime(10, 10)
	fmt.Println("Estimated brew time:", brewTime)
}
