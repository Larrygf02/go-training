package main
import (
	"fmt"
)

func main() {
	/* x := 2
	fmt.Println(x)
	{
		y:= 2
		fmt.Println(y)
	} */
	//fmt.Println(y): retornara un error porque no esta dentro del alcance de la variable
	pos := adder()
	fmt.Println(pos(2))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Println(x)
		sum += x
		return sum
	}
}

