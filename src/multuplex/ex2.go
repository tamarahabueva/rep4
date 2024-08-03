package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Интервал очистки кольцевого буфера
const bufferDrainInterval time.Duration = 30 * time.Second

// Размер кольцевого буфера
const bufferSize int = 10

// RingIntBuffer - кольцевой буфер целых чисел
type RingIntBuffer struct {
	array []int // более низкоуровневое хранилище нашего
	// буфера
	pos  int        // текущая позиция кольцевого буфера
	size int        // общий размер буфера
	m    sync.Mutex // мьютекс для потокобезопасного доступа к
	// буферу.
	// Исключительный доступ нужен,
	// так так одновременно может быть вызваны
	// методы Get и Push,
	// первый - когда настало время вывести
	// содержимое буфера и очистить его,
	// второй - когда пользователь ввел новое
	// число, оба события обрабатываются разными
	// горутинами.

}

// NewRingIntBuffer - создание нового буфера целых чисел
func NewRingIntBuffer(size int) *RingIntBuffer {
	return &RingIntBuffer{make([]int, size), -1, size, sync.Mutex{}}
}

// Push добавление нового элемента в конец буфера
// При попытке добавления нового элемента в заполненный буфер
// самое старое значение затирается
func (r *RingIntBuffer) Push(el int) {
	r.m.Lock()
	defer r.m.Unlock()
	if r.pos == r.size-1 {
		// Сдвигаем все элементы буфера
		// на одну позицию в сторону начала
		for i := 1; i <= r.size-1; i++ {
			r.array[i-1] = r.array[i]
		}
		r.array[r.pos] = el
	} else {
		r.pos++
		r.array[r.pos] = el
	}
}

// Get - получение всех элементов буфера и его последующая очистка
func (r *RingIntBuffer) Get() []int {
	if r.pos < 0 {
		return nil
	}
	r.m.Lock()
	defer r.m.Unlock()
	var output []int = r.array[:r.pos+1]
	// Виртуальная очистка нашего буфера
	r.pos = -1
	return output
}

// StageInt - Стадия конвейера, обрабатывающая целые числа
type StageInt func(<-chan bool, <-chan int) <-chan int

// PipeLineInt - Пайплайн обработки целых чисел
type PipeLineInt struct {
	stages []StageInt
	done   <-chan bool
}

// NewPipelineInt - Создание пайплайна обработки целых чисел
func NewPipelineInt(done <-chan bool, stages ...StageInt) *PipeLineInt {
	return &PipeLineInt{done: done, stages: stages}
}

// Run - Запуск пайплайна обработки целых чисел
// source - источник данных для конвейера
func (p *PipeLineInt) Run(source <-chan int) <-chan int {
	var c <-chan int = source
	for index := range p.stages {
		c = p.runStageInt(p.stages[index], c)
	}
	return c
}

// runStageInt - запуск отдельной стадии конвейера
func (p *PipeLineInt) runStageInt(stage StageInt, sourceChan <-chan int) <-chan int {
	return stage(p.done, sourceChan)
}
func main() {
	// источник данных
	dataSource := func() (<-chan int, <-chan bool) {
		c := make(chan int)
		done := make(chan bool)
		go func() {
			defer close(done)
			scanner := bufio.NewScanner(os.Stdin)
			var data string
			for {
				scanner.Scan()
				data = scanner.Text()
				if strings.EqualFold(data, "exit") {
					fmt.Println("Программа завершила работу!")
					return
				}
				i, err := strconv.Atoi(data)
				if err != nil {
					fmt.Println("Программа обрабатывает только целые числа!")
					continue
				}
				c <- i
			}
		}()
		return c, done
	}
	// стадия, фильтрующая отрицательные числа
	negativeFilterStageInt := func(done <-chan bool, c <-chan int) <-chan int {
		convertedIntChan := make(chan int)
		go func() {
			for {
				select {
				case data := <-c:
					if data > 0 {
						select {
						case convertedIntChan <- data:
						case <-done:
							return
						}
					}
				case <-done:
					return
				}
			}
		}()
		return convertedIntChan
	}
	// стадия, фильтрующая числа, не кратные 3
	specialFilterStageInt := func(done <-chan bool, c <-chan int) <-chan int {
		filteredIntChan := make(chan int)
		go func() {
			for {
				select {
				case data := <-c:
					if data != 0 && data%3 == 0 {
						select {
						case filteredIntChan <- data:
						case <-done:
							return
						}
					}
				case <-done:
					return
				}
			}
		}()
		return filteredIntChan
	}
	// стадия буферизации
	bufferStageInt := func(done <-chan bool, c <-chan int) <-chan int {
		bufferedIntChan := make(chan int)
		buffer := NewRingIntBuffer(bufferSize)
		go func() {
			for {
				select {
				case data := <-c:
					buffer.Push(data)
				case <-done:
					return
				}
			}
		}()
		// В этой стадии есть вспомогательная горутина,
		// выполняющая просмотр буфера с заданным интервалом
		// времени -
		// bufferDrainInterval
		go func() {
			for {
				select {
				case <-time.After(bufferDrainInterval):
					bufferData := buffer.Get()
					// Если в кольцевом буфере что-то есть -
					// выводим
					// содержимое построчно
					if bufferData != nil {
						for _, data := range bufferData {
							select {
							case bufferedIntChan <- data:
							case <-done:
								return
							}
						}
					}
				case <-done:
					return
				}
			}
		}()
		return bufferedIntChan
	}
	// Потребитель данных от пайплайна
	consumer := func(done <-chan bool, c <-chan int) {
		for {
			select {
			case data := <-c:
				fmt.Printf("Обработаны данные: %d\n", data)
			case <-done:
				return
			}
		}
	}
	// Запускаем наш воображаемый источник данных,
	// он же ответственен за сигнализирование о том,
	// что он завершил работу
	source, done := dataSource()
	// Создаем пайплайн, передаем ему специальный канал,
	// синхронизирующий завершение работы пайплайна,
	// а также передаем ему все стадии
	pipeline := NewPipelineInt(done, negativeFilterStageInt, specialFilterStageInt, bufferStageInt)
	consumer(done, pipeline.Run(source))
}
