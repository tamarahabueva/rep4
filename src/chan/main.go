package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gor1(c chan int) {
	for i := 0; i < 10; i++ {
		i := rand.Intn(10)
		c <- i
	}

}

func gor2(c chan int) {
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		i := rand.Intn(10)
		c <- i
	}
}

func main() {

	chan1 := make(chan int)
	chan2 := make(chan int)

	go gor1(chan1)
	go gor2(chan2)

	time.Sleep(1 * time.Second)

	select {
	case res := <-chan1:
		fmt.Println("Целые числа из горутины 1:", res)
	case res := <-chan2:
		fmt.Println("Целые числа из горутины 2:", res)
	default:
		fmt.Println("дефолт")
	}
}
