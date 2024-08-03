package main

import (
	"fmt"
)

func main() {
	// Небольшая обертка над множеством чисел, которая направляет их
	// в канал
	// и возвращает его наружному коду
	init := func(done <-chan int, integers ...int) <-chan int {
		output := make(chan int)
		go func() {
			defer close(output)
			for _, i := range integers {
				select {
				case <-done:
					return
				case output <- i:
				}
			}
		}()
		return output
	}
	// Потокобезопасная версия стадии конвейера, осуществляющая
	// выполнения арифметического произведения данных на заданное
	// число
	multiply := func(done <-chan int, input <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for {
				select {
				case <-done:
					return
				case i, isChannelOpen := <-input:
					if !isChannelOpen {
						return
					}
					select {
					case multipliedStream <- i * multiplier:
					case <-done:
						return
					}
				}
			}

		}()
		return multipliedStream
	}
	// Потокобезопасная версия стадии конвейера, осуществляющая
	// выполнения арифметического сложения данных на заданное число
	add := func(done <-chan int, input <-chan int, additive int) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for {
				select {
				case <-done:
					return
				case i, isChannelOpen := <-input:
					if !isChannelOpen {
						return
					}
					select {
					case addedStream <- i + additive:
					case <-done:
						return
					}
				}
			}
		}()
		return addedStream
	}
	// Этот канал используется для централизованной остановки
	// конвейера
	// Так как ни одна стадия не должна знать, когда конвейер будет
	// остановлен,
	// это не задача стадии, как вы понимаете,
	// Это лучше делать вот так - централизованно
	done := make(chan int)
	defer close(done)
	intStream := init(done, 1, 2, 3, 4)
	// Одно из преимуществ использования каналов видно здесь -
	// каналы итерируемы, благодаря этому мы можем комбинировать
	// пайплайн как хотим,
	// а сама целостность пайплайна от этого не меняется
	pipeline := multiply(done, add(done, multiply(done, intStream, 4), 1), 2)
	for v := range pipeline {
		fmt.Println(v)
	}
}
