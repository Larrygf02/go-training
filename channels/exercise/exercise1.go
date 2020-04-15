package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)

	text := make(chan string)
	go func() { text <- "hello text" }()
	mtext := <-text
	fmt.Println(mtext)

	numbers := make(chan int)
	go func() { numbers <- 3 }()
	num := <-numbers
	fmt.Println(num)
}
