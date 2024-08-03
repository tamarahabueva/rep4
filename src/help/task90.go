package main

import "fmt"

func IsPalindrome(n int) bool {
	a, b := 0, n
	for b > 0 {
		a *= 10
		a += b % 10
		b /= 10
	}
	return a == n
}
func main() {
	val := IsPalindrome(545)
	fmt.Println(val)
}
