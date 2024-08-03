package main

import "fmt"

func main() {
	n := 10
	sl := make([]float64, 0, 10)
	fmt.Println(sl)
	for i := 0; i < n; i++ {
		sl = append(sl, float64(i))
	}
	fmt.Println(sl)
}
