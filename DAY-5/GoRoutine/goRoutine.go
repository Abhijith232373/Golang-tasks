package main

import (
	"fmt"
	"time"
)

func odd() {
	for i:= 1;i<=9; i++ {
		if i%2 != 0 {
			fmt.Println(i)
			time.Sleep(100*time.Millisecond)
		}
	}
}

func  even(){
for i:=2;i<=10;i++{
	if i%2==0{
		fmt.Println(i)
		time.Sleep(100*time.Millisecond)
	}
}
}
func main() {
	go odd()
	go even()
time.Sleep(2*time.Second)
fmt.Println("done")
}