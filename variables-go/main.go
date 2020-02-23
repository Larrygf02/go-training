package main

import "fmt"

func main()  {
	identifiers()
}

func identifiers()  {
	/* declarado y asignado*/
	number := 42
	fmt.Println(number)
	number = 54
	fmt.Println(number)

	sum := 20 + 32
	fmt.Println(sum)
}
