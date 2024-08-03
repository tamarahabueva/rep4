package main

import (
	"fmt"
)

func main() {
	val, err := findMax([]int{1, 2, 3, 2, 4, 2, 5})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func findMax(array []int) (max int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found max in empty slice")
	}

	max = array[0]
	for _, val := range array[1:] {
		if val > max {
			max = val
		}
	}

	return max, nil
}
