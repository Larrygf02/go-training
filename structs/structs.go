package main
import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

func main() {
	p1 := person{
		first: "James",
		last: "Bond",
	}
	p2 := person{
		first: "Miss",
		last: "Moneypenny",
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("Hello, play")
}