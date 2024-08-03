package main

import "fmt"

func main() {
	// Ваш код помещайте после этой строки и до символа }
	const (
		a = iota * 2
		b
		c
	)
	fmt.Println(a, b, c)

}
