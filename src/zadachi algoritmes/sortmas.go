package main

import "fmt"

func main() {
	a := []int{5, 7, 3, 2, 1, 4, 6, 9}
	fmt.Println(checkSliceIsSorted(a))
}

func checkSliceIsSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i+1] {
			return false
		}
	}
	return true

}
