package main

import (
	"fmt"
	"time"
)

func main() {
	go gorut()
	time.Sleep(2 * time.Second)
}
func gorut() {
	fmt.Println("hi")
}
