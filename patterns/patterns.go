package main

import (
	"fmt"

	"github.com/larrygf02/patterns/client_one"
	"github.com/larrygf02/patterns/client_two"
	"github.com/larrygf02/patterns/singlenton"
)

func main() {
	fmt.Println("Hello world")
	client_one.IncrementAge()
	client_two.IncrementAge()

	p := singlenton.GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}
