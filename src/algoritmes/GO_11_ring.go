package main

import (
	"fmt"
)

// IntRing - кольцевой массив элементов, содержащих целые числа
type IntRing struct {
	size   int    // общий размер массива
	data   []*int // непосредственно контейнер, хранящий элементы массива
	Start  int    // позиция начала кольцевого массива
	End    int    // позиция конца кольцевого массива
	isFull bool   // флаг, указывающий, что кольцевой массив полон
}

// NewIntRing - конструктор кольцевого массива
func NewIntRing(size int, start int) (*IntRing, error) {
	if start >= size {
		return nil, fmt.Errorf("Стартовая позиция <%d> не соответствует размеру кольцевого массива <%d>", start, size)
	}
	return &IntRing{size, make([]*int, size, size), start, start, false}, nil
}

// Size - получение общего размера кольцевого массива
func (r IntRing) Size() int {
	return r.size
}

// IsEmpty - проверка, пуст ли кольцевой массив
func (r *IntRing) IsEmpty() bool {
	actualDataArray := r.getContinuousArray()
	for _, el := range actualDataArray {
		if el != nil {
			return false
		}
	}
	return true
}

// getContinuousArray - получает данные кольцевого массива в виде непрерывной области памяти
// вспомогательный метод
func (r IntRing) getContinuousArray() []*int {
	continuouslDataArray := []*int{}
	if r.Start == r.End && r.IsFull() {
		continuouslDataArray = append(r.data[r.Start:], r.data[:r.End]...)
	} else if r.Start < r.End {
		continuouslDataArray = r.data[r.Start:r.End]
	} else if r.Start > r.End {
		continuouslDataArray = append(r.data[r.Start:], r.data[:r.End]...)
	}
	return continuouslDataArray
}

// IsFull - достигнут ди конец
func (r *IntRing) IsFull() bool {
	return r.isFull
}

// Read - чтение элемента из кольцевого массива
func (r *IntRing) Read() (int, error) {
	if !r.IsEmpty() {
		var el *int
		continuouslDataArray := r.getContinuousArray()
		// перебираем данные массива как непрерывную память, для согласования индексов
		// при переборе одновременно наращиваем поле индекса начального элемента кольцевого массива r.Start
		for i := 0; el == nil && i < len(continuouslDataArray); i, r.Start = i+1, r.Start+1 {
			el = r.data[r.Start]
		}
		// Если до этого массив был заполнен после успешного чтения сбрасываем этот флаг
		if r.IsFull() {
			r.isFull = false
		}
		return *el, nil
	}
	return 0, fmt.Errorf("Не удалось произвести чтение из кольцевого массива: массив пуст!")
}

// Write - добавление одного элемента в кольцевой массив
func (r *IntRing) Write(v int) error {
	if !r.IsFull() {
		r.data[r.End] = &v
		if r.End < r.Size()-1 {
			r.End++
		} else {
			// Если конец массива ушел за границу - наращиваем массив за счет начала
			r.End = 0
		}
		// Если конец массива стал указаывать на его начало
		// после операции записи - ставим флаг, что он полон
		if r.End == r.Start {
			r.isFull = true
		}
		return nil
	}
	return fmt.Errorf("Не удалось произвести запись в кольцевой массив: массив полон!")
}

// RemoveByIndex - удаление элемента кольцевого массива
func (r *IntRing) RemoveByIndex(index int) error {
	if index < 0 || index >= r.size {
		return fmt.Errorf("<%d> - Неверный индекс удаляемого элемента кольцевого массива", index)
	}
	if r.data[index] != nil {
		r.data[index] = nil
	}
	return nil
}

// Print - печать содержимого кольцевого массива
func (r IntRing) Print() {
	if r.IsEmpty() {
		fmt.Println("Кольцевой массив пуст!")
	} else {
		dataForPrint := r.getContinuousArray()
		for _, el := range dataForPrint {
			if el != nil {
				fmt.Printf("%d\t", *el)
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	ring, err := NewIntRing(5, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	ring.Print()
}
