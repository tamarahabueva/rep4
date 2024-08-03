package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var n int
	fmt.Print("Ведите длительность работы программы в секундах: ")
	fmt.Fscan(os.Stdin, &n)

	counter := 0

	go func() {
		for {
			time.Sleep(time.Second)
			counter++
		}
	}()

	go func() {
		for {
			fmt.Printf("%v - %v\n", time.Now().Format("15:04:05.00"), counter)
			time.Sleep(time.Millisecond * time.Duration(200))
		}
	}()

	time.Sleep(time.Second * time.Duration(n))
}
