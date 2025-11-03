package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var counter int


func incremnt(wg *sync.WaitGroup){
	defer wg.Done()

	for i:=0;i<5;i++{
		mu.Lock()
		counter++
		fmt.Println("count:",counter)
		mu.Unlock()
		time.Sleep(100*time.Millisecond)
	}
}

func main() {
var wg sync.WaitGroup
wg.Add(2)
go incremnt(&wg)
go incremnt(&wg)
wg.Wait()
fmt.Println("final:",counter)
}