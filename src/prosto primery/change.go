package main

import "fmt"

func changeSlice(sl []float64) {
	sl[4] = 100
}

func main() {
	sl := []float64{1, 2, 3, 4, 5}
	changeSlice(sl)
	fmt.Println(sl)
}
