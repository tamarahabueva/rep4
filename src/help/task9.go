package main

import "fmt"

func main() {

	var originalNumber int = 303
	var reverseNumber int = 0
	var temp = originalNumber

	for temp > 0 {
		var remainder = temp % 10
		reverseNumber = (reverseNumber * 10) + remainder
		temp = temp / 10
		fmt.Println(reverseNumber)
	}
}

if originalNumber == reverseNumber {
	fmt.Println("The given number is Palindrome")
} else {
	fmt.Println("The given number is NOT a Palindrome")
}
