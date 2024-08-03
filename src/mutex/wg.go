package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("Горутина", i, "начала работу")
			// выполнение задачи
			fmt.Println("Горутина", i, "завершила работу")
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Все горутины завершили работу")
}
