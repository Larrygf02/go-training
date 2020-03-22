package main
import (
	"fmt"
)

func main() {
	numbers := []int{1,2,3,4,5,6}
	sum := sum("", numbers...)
	fmt.Println("The total is: ",sum)
}

func sum(type_sum string, numbers ...int) int {
	fmt.Println(numbers)
	fmt.Println(type_sum)
	fmt.Printf("%T\n", numbers)
	sum := 0
	if type_sum == "pares" {
		for _, v := range numbers {
			if v % 2 == 0 {
				sum += v
			}
		}
		return sum;
	}
	for _, v := range numbers {
		sum += v
	}
	return sum;
}