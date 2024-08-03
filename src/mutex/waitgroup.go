package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // в группе две горутины
	work := func(id int) {
		defer wg.Done()
		fmt.Printf("Горутина %d начала выполнение \n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Горутина %d завершила выполнение \n", id)
	}

	// вызываем горутины
	go work(1)
	go work(2)

	wg.Wait() // ожидаем завершения обоих горутин
	fmt.Println("Горутины завершили выполнение")
}
