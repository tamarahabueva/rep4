package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	const routinesN = 5
	const printTimes = 10
	wg.Add(routinesN)
	for routine := 1; routine <= routinesN; routine++ {
		go func(routine int) {
			for i := 0; i < printTimes; i++ {
				fmt.Println(routine)
			}
			wg.Done()
		}(routine)
	}
	wg.Wait()
}
