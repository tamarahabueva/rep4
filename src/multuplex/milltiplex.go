package main

import (
	"fmt"
	"sync"
)

// Число сообщений от каждого источника
const messagesAmountPerGoroutine int = 5

func main() {
	dataSourceChans := make([]chan int, 0)
	multiplexingFunc := func(channels ...chan int) <-chan int {
		var wg sync.WaitGroup
		// Общий канал, в который будут попадать сообщения от всех
		// источников
		// Именно его мы и вернём из этой функции для употребления
		// внешним кодом
		multiplexedChan := make(chan int)
		multiplex := func(c <-chan int) {
			defer wg.Done()
			for i := range c {
				// Если поступило сообщение из одного из
				// каналов-источников,
				// перенаправляем его в общий канал
				multiplexedChan <- i
			}
		}
		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}
		// Запускаем горутину, которая закроет канал после того,
		// как в закрывающий канал поступит сигнал о прекращении работы
		// всех
		go func() {
			wg.Wait()
			close(multiplexedChan)
		}()
		return multiplexedChan
	}
	// Горутина - источник данных
	// Функция создаёт свой собственный канал
	// и посылает в него пять сообщений
	startDataSource := func(start int) chan int {
		c := make(chan int)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := start; i < start+messagesAmountPerGoroutine; i++ {
				c <- i
			}
		}()
		go func() {
			wg.Wait()
			close(c)
		}()
		return c
	}
	// Запускаем пять источников
	// Канал от каждого сохраняем в наш специальный буфер
	for i := messagesAmountPerGoroutine; i <= 20; i += messagesAmountPerGoroutine {
		dataSourceChans = append(dataSourceChans, startDataSource(i))
	}
	// Уплотняем каналы
	c := multiplexingFunc(dataSourceChans...)

	// Централизованно получаем сообщения от всех нужным нам источников
	// данных
	for data := range c {
		fmt.Println(data)
	}
}
