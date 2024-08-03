package main

import "fmt"

func main() {
	const (
		january = iota + 1
		february
		march
		april
		may
	)
	fmt.Println(january, february, march, april, may)
}
