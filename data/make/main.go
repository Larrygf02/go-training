package main

import "fmt"

func main() {
	name := make([]string, 3)
	fmt.Println("names: ", name)
	name[0] = "Grace"
	name[1] = "Raul"
	name[2] = "Maria"
	name = append(name, "Yeli", "Ernesto")
	fmt.Println("names", name)
	fmt.Println("len", len(name))
	// copia name
	other_name := make([]string, len(name))
	copy(other_name, name)
	fmt.Println("Copia del arreglo name", other_name)
}
