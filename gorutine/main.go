package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	fmt.Println("Go rutine")
	// estas dos funciones se ejecutan en paralelo
	wg.Add(1)
	wg.Add(1)
	go compute(5)
	go compute(5)
	wg.Wait()
}
