package main

import "fmt"

func main() {
	val := isPalindrome(303)
	fmt.Println(val)
}
func isPalindrome(n int) bool {
	a, b := 0, n
	for b > 0 {
		a *= 10
		a += b % 10
		b /= 10
	}
	return a == n
}
