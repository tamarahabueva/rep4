package main

import "fmt"

func main() {
	var a float64
	var b float64
	var c string
	fmt.Scan(&a, &b)
	switch c {
	case "+":
		fmt.Println("Result sum", a+b)
	case "-":
		fmt.Println("Result difference", a-b)
	case "*":
		fmt.Println("Result multiplication", a-b)
	case "/":
		fmt.Println("Result division", a/b)
	default:
		fmt.Println("ОШИБКА", a/0, b/0)
	}
}
