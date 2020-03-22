package main
import (
	"fmt"
)

type person struct {
	name string
	age int
	salary int
}

func (p person) eat(food string) {
	fmt.Println("I eat",food)
}

func (p person) hello() {
	fmt.Println("Hello my name is", p.name)
}

type human interface {
	speak()
}

func main() {
	newperson := person{
		"Raul",
		26,
		3000,
	}
	newperson.hello()
	newperson.eat("Ceviche")
}