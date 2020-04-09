package main

import "fmt"

func product(nums ...int) int {
	fmt.Print(nums, " ")
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total
}

func main() {
	fmt.Println(product(1, 2))
	fmt.Println(product(1, 2, 3))
}
