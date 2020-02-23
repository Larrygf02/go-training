package main

import (
"fmt"
)

var x = 42
var y = "James Bond"
var z = true
func main() {
	fmt.Println("Hello, playground")
	print()
}

func print()  {
	s := fmt.Sprintf("%d, %s, %t", x, y, z)
	fmt.Println(s)
}
