package main
import (
	"fmt"
)

func main() {
	// argument Doe
	hello("Doe")
	name := getName("Michael")
	fmt.Println(name)
	age := getAges(33)
	fmt.Println(age)
	items(1,2,3,4)
}

// parameters
func hello(s string) {
	fmt.Println("Hello", s)
}

func getName(name string) string{
	return fmt.Sprint("My name is ", name)
}

func getAges(age int) int {
	return age;
}

func items(x ...int) {
	fmt.Println(x)
	sum := 0
	for _, value := range x {
		sum += value
	}
	fmt.Println("Sum is, ", sum)
}