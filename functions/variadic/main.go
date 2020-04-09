package main

import "fmt"

func product(nums ...int) {
	fmt.Print(nums, " ")
	total := 1
	for _, num := range nums {
		total *= num
	}
	fmt.Println(total)

}

func main() {
	product(1, 2)
	product(1, 2, 3)
}
