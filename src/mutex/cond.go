package main

import (
	"fmt"
	"sync"
)

var cond sync.Cond
var data []string
var loaded bool

func loadData() {
	fmt.Println("Загрузка данных")
	// загрузка данных
	data = []string{"данные1", "данные2", "данные3"}
	loaded = true
	cond.Signal()
}

func getData() []string {
	cond.L.Lock()
	for !loaded {
		cond.Wait()
	}
	cond.L.Unlock()
	return data
}

func main() {
	cond.L = &sync.Mutex{}
	go loadData()
	fmt.Println(getData())
}
