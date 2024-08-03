package main 
import ("fmt")
func main () {
	
	res := make(chan int)

go func(chan)
go func2(chan)
}
	  func(c chan int) {
	for i := 0; i < 10; i++ {
		c<-i
	}
	}
func2 (c chan int) {
	num <- c
	fmt.Println(num)
}

