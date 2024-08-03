package main

import (
	"errors"
	"fmt"
)

func main() {
	i, err := someFunc(0)
	fmt.Println(i, err)

	i, err = someFunc(10)
	fmt.Println(i, err)
}

func someFunc(i int) (int, error) {
	if i == 0 {
		return 0, errors.New("got 0")
	}

	return i, nil
}
