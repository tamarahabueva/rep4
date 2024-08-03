package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Создаем два кольца, каждый размером по 3 элемента
	ring1 := ring.New(3)
	ring2 := ring.New(3)

	// Получаем длины каждого из колец
	ring1Len := ring1.Len()
	ring2Len := ring2.Len()

	// Заполняем первое кольцо нулями
	for i := 0; i < ring1Len; i++ {
		ring1.Value = 0
		ring1 = ring1.Next()
	}

	// Заполняем второе кольцо единичками
	for j := 0; j < ring2Len; j++ {
		ring2.Value = 1
		ring2 = ring2.Next()
	}

	// Соединяем второе кольцо с первым вот таким образом
	ring1.Next().Link(ring2)
	// Так как мы соединили два кольца -
	// размер нового кольца равен сумме размеров исходных
	unionLen := ring1Len + ring2Len

	// Проверяем содержимое нового кольца
	for i := 0; i < unionLen; i++ {
		fmt.Println(ring1.Value)
		ring1 = ring1.Next()
	}
}
