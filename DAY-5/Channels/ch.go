package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("this is Producerss:",i)
		ch<-i
		time.Sleep(time.Second)
	}
	close(ch)
}

func cutomer(ch chan int){
	for item:=range ch{
		fmt.Println("this is cutomers:",item)
		time.Sleep(3*time.Millisecond)
	}
}
func main() {
	ch := make(chan int, 2)
	go producer(ch)
	cutomer(ch)

}