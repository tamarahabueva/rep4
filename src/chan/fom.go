package main

import (
	"fmt"
	"math/rand"
	"time"
)

func serv1(c chan int) {
	for i := 0; i < 10; i++ {
		i := rand.Intn(10)
		c <- i
	}
}
func serv2(c chan int) {
	for i := 0; i < 10; i++ {
		i := rand.Intn(10)
		c <- i
	}

}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go serv1(chan1)
	go serv2(chan2)
	go func() {
		for {
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		select {
		case res := <-chan1:
			fmt.Println("Получено из канала 1:", res)
		case res := <-chan2:
			fmt.Println("Получено из канала 2:", res)
		}
	}
}
