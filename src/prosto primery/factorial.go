package main

import "fmt"

func main() {
	fmt.Println(recursive(5))
}

func recursive(number int) int {
	if number > 1 {
		return number * recursive(number-1)
	}
	return 1
}
