package main

import "fmt"

func main() {
	// Untyped constant with integer value
	const rewardPoints = 10
	fmt.Printf("Default type of reward points is %T\n", rewardPoints)

	var totalRewardPoints float64 = 150.3

	// Adding untyped constant to a float64 - valid
	totalRewardPoints += rewardPoints

	fmt.Printf("Updated type of total reward points is %.2f\n", totalRewardPoints)

}
