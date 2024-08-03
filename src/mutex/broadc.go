package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

// Количество слушающих горутин
const amountOfGoroutines int = 5

// Текстовая команда, набираемая в консоли, для выхода из программы
const quitCommand string = "quit"

func main() {
	var message string
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	// Для ожидания запуска слушающих горутин
	var startWg sync.WaitGroup
	// Для ожидания завершения работы всех запущенных слушающих
	// горутин
	var closeWg sync.WaitGroup
	startWg.Add(amountOfGoroutines)
	closeWg.Add(amountOfGoroutines)
	c := sync.NewCond(&sync.Mutex{})
	for i := 1; i <= amountOfGoroutines; i++ {
		go func(id int) {
			defer closeWg.Done()
			var outPutMessage string
			// Уведомляем главную горутину, что очередная слушающая
			// горутина запущена
			startWg.Done()
			for {
				c.L.Lock()
				c.Wait()
				c.L.Unlock()
				// Сравниваем введённую строку с командой выхода, не
				// учитывая регистр символов
				if strings.EqualFold(message, quitCommand) {
					return
				}
				if utf8.RuneCountInString(message) == 0 {
					outPutMessage = fmt.Sprintf("Горутина №%d обработала событие: Осуществлен ввод пустой строки", id)
				} else {
					outPutMessage = fmt.Sprintf("Горутина №%d обработала событие: Осуществлен ввод строки: \"%s\"", id, message)
				}
				fmt.Println(outPutMessage)
			}

		}(i)
	}
	// Ожидаем запуска всех слушающих горутин
	startWg.Wait()

	for {
		scanner.Scan()
		message = scanner.Text()
		fmt.Println("------------------------------------------------------------")
		c.Broadcast()
		// Сравниваем введённую строку с командой выхода, не учитывая
		// регистр символов
		if strings.EqualFold(message, quitCommand) {
			// Ждём завершение работы всех запущенных горутин
			closeWg.Wait()
			fmt.Println("Выход из программы")
			return
		}
	}
}
