package main

import (
	"fmt"
	"time"
)

var (
	semaphoreCounterSize = 5
	semaphoreChan        = make(chan int, semaphoreCounterSize)
)

// Обработка некой информации
func operate(item int) {
	semaphoreChan <- 0
	go func() {
		defer func() {
			<-semaphoreChan
		}()
		fmt.Println(item)
		time.Sleep(1 * time.Second)
	}()
}

func main() {
	// Запускаем десять обработчиков, которые будут обращаться к общему ресурсу,
	// защищаемому нашим семафором
	for i := 0; i < 10; i++ {
		operate(i)
	}
}
