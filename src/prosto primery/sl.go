package main

import "fmt"

func main() {

	var arr [5]int
	sl := arr[:]
	fmt.Println(sl)
	sl = arr
	fmt.Println(sl)
}
