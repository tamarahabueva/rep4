package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println("Hello, goroutine!")
			time.Sleep(time.Second)
		}()
	}
	time.Sleep(time.Second)
}
