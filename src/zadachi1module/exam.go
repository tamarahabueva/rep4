package main

import "fmt"

func main() {
	loopFive()
}

func loopFive() {
	for i := 0; i < 25; i++ {
		fmt.Print(i)
		if i == 5 {
			// Stop function at i == 5
			break
		}
	}
	fmt.Println("This line will not execute.")
}
