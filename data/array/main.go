package main

import "fmt"

func main() {
	slices()
}

func arrayMake() {
	fmt.Println("Array Make")
}

func slices() {
	s := []int{2, 3, 4, 5, 6, 7}
	s = s[1:4]
	fmt.Println(s)
}

func intro() {
	var s []int
	fmt.Println(s)
	// add a element
	s = append(s, 0)
	fmt.Println(s)
	// add more a element
	s = append(s, 1, 2)
	y := []int{3, 3}
	s = append(s, y...)
	fmt.Println(s)

	// Eliminar un elemento del indice i
	// slice = append(slice[:i], slice[i+1]...)
	slice := append(s[:0], s[2:]...)
	fmt.Println(slice)
}
