package main

import "fmt"

func main() {
	name := make([]string, 3)
	fmt.Println("names: ", name)
	name[0] = "Grace"
	name[1] = "Raul"
	name[2] = "Maria"
	fmt.Println("names", name)
}
