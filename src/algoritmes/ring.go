package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Создаем кольцо из пяти элементов
	r := ring.New(5)
	// Получаем длину
	n := r.Len()
	// Инициализируем элементы кольца значениями
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	// Перемещаемся вперёд по кольцу и печатаем значение каждого
	// элемента
	for j := 0; j < n; j++ {
		fmt.Println(r.Value)
		r = r.Next()
	}
	// Перемещаемся назад по кольцу и печатаем значение каждого
	// элемента
	for k := n; k > 0; k-- {
		r = r.Prev()
		fmt.Println(r.Value)
	}
}
