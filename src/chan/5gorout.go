package main
 
import (
    "fmt"
    "sync"
    "time"
)
 
// Количество создаваемых горутин
const goroutineAmount int = 5
 
func main() {
    begin := make(chan interface{})
    var wg sync.WaitGroup
    // Цикл запуска пяти горутин
    for i := 0; i < goroutineAmount; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            <-begin
            fmt.Printf ("Горутина №%d получила сигнал о закрытии и 
                        завершила свою работу\n", i)
        }(i)
    }
    fmt.Println("Оповещение и разблокировка горутин произойдёт через 5 
                 секунд...")
    // Некий другой продолжительный по времени выполнения код,
    // вместо которого - просто блокировка на 5 секунд
    time.Sleep(time.Second * 3)
    close(begin)
    wg.Wait()
}