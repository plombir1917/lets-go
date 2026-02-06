package main

import "fmt"

func slices() {
	slice := []int{}
	i := 1
	for i <= 100 {
		slice = append(slice, i)
		i++
	}
	fmt.Println(slice)

	slice = append(slice[:10], slice[89:]...)
	fmt.Println(slice)
}

func main() {
	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
	workDaysSlice := weekTempArr[:5]
	weekendDaysSlice := weekTempArr[5:]
	fromTuesdayToThursdaySlice := weekTempArr[1:4]
	weekTempSlice := weekTempArr[:]

	weekTempSlice[1] = 100
	fmt.Println(workDaysSlice)
	fmt.Println(weekendDaysSlice)
	fmt.Println(fromTuesdayToThursdaySlice)
	fmt.Println(weekTempSlice)
	slices()
}
