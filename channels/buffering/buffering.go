package main

import "fmt"

func main() {
	messages := make(chan string, 3)
	messages <- "grace"
	messages <- "y"
	messages <- "raul"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
