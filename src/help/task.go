package main

import "fmt"

func main() {
	var n int = 565
	var a int = 0
	var b int = n

	a *= 10
	a += b % 10
	b /= 10
	fmt.Println(a)
	fmt.Println(b)
}
