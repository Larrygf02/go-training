package main
import (
	"fmt"
)

type person struct {
	name string
	age int
	salary int
}

func (p person) hello() {
	fmt.Println("Hello my name is ", p.name)
}

func main() {
	newperson := person{
		"Raul",
		26,
		3000,
	}
	newperson.hello()
}