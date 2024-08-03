package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100
	}

	insertionSort(ar)

	fmt.Println(ar)
}

func insertionSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	for i := 1; i < len(ar); i++ {
		for j := i; j > 0 && ar[j-1] > ar[j]; j-- {
			ar[j-1], ar[j] = ar[j], ar[j-1]
		}
	}
}
