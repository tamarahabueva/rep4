package main

import "fmt"

func main() {
	val, err := findMostOftenRepeated([]int{1, 2, 3, 4, 3})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}
func findMostOftenRepeated(array []int) (mostOften int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	var maxIndex, maxCount = 0, 0
	for i, number := range array {
		currentCount := 0
		for _, numberToCompare := range array {
			if number == numberToCompare {
				currentCount++
			}
		}

		if currentCount > maxCount {
			maxIndex = i
			maxCount = currentCount
		}
	}

	return array[maxIndex], nil
}
