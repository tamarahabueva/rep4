package main

import "fmt"

func isPalindrome(x int) bool {
	var sum int = 0
	var compare = x
	for x > 0 {
		r := x % 10
		sum = (sum * 10) + r
		x = x / 10

	}
	if compare == sum {
		return true
	}
	return false
}
func main() {
	val := isPalindrome(575)
	fmt.Println(val)
}
