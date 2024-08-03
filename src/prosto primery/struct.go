package main

import "fmt"

type Creature struct {
	Name string
}

func main() {
	c := Creature{
		Name: "Sammy the Shark",
	}
	fmt.Println(c.Name)
}