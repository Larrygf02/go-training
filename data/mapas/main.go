package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 4
	m["k2"] = 5
	fmt.Println("map", m)
	v1 := m["k1"]
	fmt.Println("item", v1)
	fmt.Println("longitud", len(m))
	delete(m, "k2")
	fmt.Println("map", m)
	m["k1"] = 6
	fmt.Println("map", m)
	fmt.Println("item", v1)
}
