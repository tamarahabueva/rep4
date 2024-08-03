package main

import "fmt"

func main() {
	intCh := make(chan int)

	go func() {
		fmt.Println("Go routine starts")
		intCh <- 5 // блокировка, пока данные не будут получены функцией main
	}()
	fmt.Println(<-intCh) // получение данных из канала
	fmt.Println("The End")
}
