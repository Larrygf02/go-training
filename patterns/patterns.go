package main

import (
	"fmt"
	"sync"

	"github.com/larrygf02/patterns/client_one"
	"github.com/larrygf02/patterns/client_two"
	"github.com/larrygf02/patterns/singlenton"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client_one.IncrementAge()
		}()
		go func() {
			defer wg.Done()
			client_two.IncrementAge()
		}()
	}
	fmt.Println("Hello world")
	wg.Wait()
	p := singlenton.GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}
