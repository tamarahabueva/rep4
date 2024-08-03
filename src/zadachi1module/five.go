package main

import "log"

// программа считает сумму чётных чисел от 0 до 50 включительно и выводит её
func main() {
	var i int
	var sum int64

	for i := 0; i < 50; i++ {
		if i%2 != 0 {
			break
		}
		sum = int64(i)
		sum += i

		log.Println(sum)
	}
}

