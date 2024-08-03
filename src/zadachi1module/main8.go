package main

import "fmt"

func Fact(n int) int {
	res := int (1)
	for i := int (1); i <= n; i++ {
		res *= i
	}
	return res
}
func main() {
	fmt.Println(Fact(n))
}
