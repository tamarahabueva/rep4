package main

import "fmt"

func fact(n int) int {
	res := int(1)
	for i := int(1); i <= n; i++ {
		res *= i
	}
	return res
}
func main() {
	val := fact(4)
	fmt.Println(val)
}
