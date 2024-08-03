package main

import (
	"fmt"
)

func main() {
	i := 4
	defer func(n int) {
		fmt.Println("Everything is fine:", n)
	}(i)
	i = 3
	fmt.Println(i)
}
