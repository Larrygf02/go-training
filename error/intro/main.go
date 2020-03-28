package main

import "fmt"

func main() {
	var answer int
	fmt.Print("Name: ")
	_, err := fmt.Scan(&answer)
	if err != nil {
		// aborta la función
		panic(err)
	}
}
