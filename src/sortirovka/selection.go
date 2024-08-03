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

	selectionSort(ar)

	fmt.Println(ar)
}

func selectionSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		var minIndex = i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}

		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
}
