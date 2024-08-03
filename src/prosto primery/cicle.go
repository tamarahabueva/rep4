package main

import "fmt"

func main() {
	number := 0

	for i := 20; i > 0; i-- {
		number = i
	}
	fmt.Println(number)
}
