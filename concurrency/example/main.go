// En este ejemplo se usara las gorutinas
package main

import "fmt"

var x = 0

func increment() {
	// seccion critica
	x = x + 1
}

func main() {
	// si ejecutams saldra distintos valores
	for i := 0; i < 1000; i++ {
		go increment()
	}
	fmt.Println("Final value of x", x)
}
