package main
import "fmt"

type person struct {
	name string
	address string
}

func (p *person) change(address string) {
	p.address = address
}
func main() {
	newperson := person{
		"Raul",
		"Av D",
	}
	fmt.Println(newperson)
	newperson.change("Av 26 de noviembre");
	fmt.Println(newperson)
}