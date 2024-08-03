package main

import (
	"fmt"
)

// ErrQueueOverFlow - ошибка переполнения очереди
var ErrQueueOverFlow = fmt.Errorf("Ошибка переполнения очереди!")

// ErrQueueEmpty - ошибка при чтении из пустой очереди
var ErrQueueEmpty = fmt.Errorf("Очередь пуста!")

// IntQueue - структура очереди
type IntQueue struct {
	data      []int
	headIndex int
	tailIndex int
	size      int
}

// NewIntQueue - конструктор очереди
func NewIntQueue(size int) *IntQueue {
	return &IntQueue{make([]int, size, size), 0, 0, 0}
}

// Size - получение размера очереди
func (q IntQueue) Size() int {
	return q.size
}

// MaxSize - получение максимального размера очереди
func (q IntQueue) MaxSize() int {
	return len(q.data)
}

// Tail - получение хвоста очереди
func (q IntQueue) Tail() (int, error) {
	if q.headIndex == q.tailIndex {
		return 0, ErrQueueEmpty
	}
	return q.data[q.tailIndex], nil
}

// Head - получение головы (начала) очереди
func (q IntQueue) Head() (int, error) {
	if q.headIndex == q.tailIndex {
		return 0, ErrQueueEmpty
	}
	return q.data[q.headIndex], nil
}

// Queue - добавление элемента в очередь
func (q *IntQueue) Queue(el int) error {
	if q.tailIndex == q.MaxSize() {
		return ErrQueueOverFlow
	}
	q.data[q.tailIndex] = el
	q.size++
	q.tailIndex++
	return nil
}

// Dequeue - извлечение элемента из очереди
func (q *IntQueue) Dequeue() (int, error) {
	if q.headIndex == q.tailIndex {
		return 0, ErrQueueEmpty
	}
	head := q.data[q.headIndex]
	for key := q.headIndex + 1; key < q.tailIndex; key++ {
		q.data[key-1] = q.data[key]
	}
	q.tailIndex--
	q.size--
	return head, nil
}

// Print - печать очереди
func (q IntQueue) Print() {
	if q.headIndex != q.tailIndex {
		for i := q.headIndex; i < q.tailIndex; i++ {
			fmt.Printf("%d\t", q.data[i])
		}
		fmt.Printf("\n")
	} else {
		fmt.Printf("%s\n", ErrQueueEmpty.Error())
	}
}
func main() {
	queue := NewIntQueue(5)
	queue.Print()
	_, err := queue.Head()
	if err != nil {
		fmt.Println(err.Error())
	}
	queue.Queue(0)
	queue.Queue(1)
	queue.Queue(2)
	fmt.Println(queue.Size())
	queue.Queue(3)
	queue.Queue(4)
	err = queue.Queue(5)
	if err != nil {
		fmt.Println(err.Error())
	}
	queue.Print()
	el, _ := queue.Dequeue()
	fmt.Println(el)
	fmt.Println(queue.Size())
	queue.Print()
}
