package main

import "fmt"

func main() {
	val, err := findMaxNegative([]int{1, 2, 3, 2, 4, 2, 5})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}
func findMaxNegative(array []int) (maxNegative int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found maxNegative in empty slice")
	}

	maxNegative = array[0]
	for _, val := range array[1:] {
		if val < maxNegative {
			maxNegative = val
		}
	}

	return maxNegative, nil
}
