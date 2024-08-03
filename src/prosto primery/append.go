package main

import "fmt"

func main() {
	sl := make([]float64, 5, 10)
	fmt.Println(sl, cap(sl), len(sl))
	sl = append(sl, 1, 2, 3, 4, 5)
	fmt.Println(sl, cap(sl), len(sl))
	sl = append(sl, 6)
	fmt.Println(sl, cap(sl), len(sl))
}
