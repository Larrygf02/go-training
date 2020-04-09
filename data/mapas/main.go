package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 4
	m["k2"] = 5
	fmt.Println("map", m)
}
