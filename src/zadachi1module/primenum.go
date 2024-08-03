package main

import "fmt"

func main() {

	for i := 1; i < 20; i++ {
		if i%1 == 0 && i%i == 0 {
			continue
		}
		i += i
		fmt.Println(i)
	}
}
