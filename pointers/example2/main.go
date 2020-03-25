package main
import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println(x)
	mutate(&x)
	fmt.Println(x)
}

func noMutate() {
	x := 0
	y := x
	y = 12
	fmt.Println(x)
	fmt.Println(y)
}

func mutate(y *int) {
	*y = 1
}