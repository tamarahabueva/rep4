package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	num := 32
	for i := 0; i < num; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer func() {
				fmt.Printf("Goroutine %d finished\n", i+1)
				wg.Done()
			}()
			var c uint64
			for j := uint64(0); j < 1<<num; j++ {
				c++
			}
		}()
	}
	wg.Wait()
}
