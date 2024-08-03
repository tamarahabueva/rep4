package main

import (
	"fmt"
)

func add(list []int, el int) []int {
	return append(list, el)
}
func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	newlist := add(a, 4)
	fmt.Println(newlist)
}
