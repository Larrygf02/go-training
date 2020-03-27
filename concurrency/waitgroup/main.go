package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func myFunc() {
	time.Sleep(1 * time.Second)
	fmt.Println("Finished GoRutine")
	wg.Done()
}

func main() {
	fmt.Println("Hello Golang")
	wg.Add(1)
	go myFunc()
	wg.Wait()
	fmt.Println("End program")
}
