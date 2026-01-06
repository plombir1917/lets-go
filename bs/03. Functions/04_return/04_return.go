package main

import "fmt"

func calculateLoyaltyPoints(amountSpent float64) int {
	loyaltyPoints := int(amountSpent) * 2
	return loyaltyPoints
}

func increaseTotalPoints(total int, newly int) int {
	return total + newly
}

func main() {
	totalPoints := 120
	newlyEarnedPoints := calculateLoyaltyPoints(10)
	fmt.Println(newlyEarnedPoints)

	totalPoints = increaseTotalPoints(totalPoints, newlyEarnedPoints)
	fmt.Println(totalPoints)
}
