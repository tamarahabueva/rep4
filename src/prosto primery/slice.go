package main

import "fmt"

func main() {
	var arr = [5]float64{1, 2, 3, 4, 5}
	sl := arr[:]
	fmt.Println(arr, sl)

	sl[4] = (100)
	fmt.Println(arr, sl)
}
