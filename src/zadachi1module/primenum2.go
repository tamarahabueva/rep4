package main

import (
	"fmt"
)

func main() {
	counter := 1
	fmt.Println(2)
	for i := 3; counter < 20; i++ {
		c := 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				c++
			}
		}
		if c == 0 {
			fmt.Println(i)
			counter++
		}
	}
}
