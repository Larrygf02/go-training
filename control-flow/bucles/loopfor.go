package main

import "fmt"

func main() {
	sum := 1
	// similar while
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}
