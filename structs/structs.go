package main
import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
	isSecret bool
	first string
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "bond",
		},
		first: "James coll",
		isSecret: false,
	}
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
	fmt.Println(sa1)
	// accede directaremente a la propiedad en lugar de sa1.person.first
	fmt.Println(sa1.first)
	// si hay colision con las propiedades
	fmt.Println(sa1.person.first)
	fmt.Println("Hello, play")
}