package main

import "fmt"

func main() {

	for i := 1; i < 100; i *= 2 {
		i += 4

		if i%2 == 0 {
			fmt.Println("чётное")
		}
	}
}
