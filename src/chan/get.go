package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 10; i++ {
		i := rand.Intn(10)
		fmt.Println(i)
	}
}
