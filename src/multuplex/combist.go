package main

import "fmt"

func main() {
	// Стадия умножения потока целых чисел на заданное число
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}
		return multipliedValues
	}
	// Стадия суммирования потока целых чисел с заданным числом
	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range multiply(add(ints, 2), 1) {
		fmt.Println(v)
	}

}
