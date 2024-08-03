package main
 
import (
    "fmt"
    "sync"
    "time"
)
 
func main() {
    c := make(chan string)
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        c <- "Сообщение"
        fmt.Println("Горутина отработала!")
    }()
    time.Sleep(5 * time.Second)
    fmt.Println(<-c)
    wg.Wait()
}
