package main
import (
	"fmt"
)
func main() {

	func (name string) {
		fmt.Println("My name is")
	}("Raul")
	hello :=func () {
		fmt.Println("Hello func anonymus")
	}
	hello()
}