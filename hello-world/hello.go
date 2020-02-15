package main

import "fmt"

func main () {
	fmt.Println("Hello World")
	for i:= 0; i < 100; i++ {
		if i % 7 == 0 {
			fmt.Println(i)
		}
	}
	foo()
	bar()
}

func foo ()  {
	fmt.Println("Hello foo")
}

func bar ()  {
	// No se puede declarar una variable sin usar
	n, _ := fmt.Println("Hello", 24, true)
	fmt.Println(n)
}