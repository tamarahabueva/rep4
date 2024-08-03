package main

import (
	"fmt"
)

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %s article was written by %s", a.Title, a.Author)
}

func Print(s fmt.Stringer) {
	fmt.Println(s)
}

func main() {
	a := Article{
		Title:  "Go for beginners",
		Author: "James Bond",
	}
	Print(a)
}
