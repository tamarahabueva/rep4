package main

import (
	"fmt"
)

// ErrEmptyStack Ошибка при извлечении данных из пустого стека
var ErrEmptyStack error = fmt.Errorf("%s", "Стек пуст")

func size(stack []int) int {
	return len(stack)
}
func push(stack *[]int, el int) {
	*stack = append(*stack, el)
}
func pop(stack *[]int) (int, error) {
	size := size(*stack)
	if size == 0 {
		return 0, ErrEmptyStack
	}
	el := (*stack)[size-1]
	*stack = (*stack)[:size-1]
	return el, nil
}
func main() {
	a := &[]int{}
	push(a, 1)
	push(a, 2)
	push(a, 3)
	el1, _ := pop(a)
	fmt.Println(el1)
	el2, _ := pop(a)
	fmt.Println(el2)
	el3, _ := pop(a)
	fmt.Println(el3)
	_, err := pop(a)
	if err != nil {
		fmt.Println(err)
	}
}
