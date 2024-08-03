package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	const k = 5
	wg.Add(k)
	for i := 1; i <= k; i++ {
		go func(i int) {
			for j := 1; j <= 10; j++ {
				fmt.Println(i)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
