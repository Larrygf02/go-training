package main

import (
	"fmt"
	"strings"
)

func DoubleChars(original string) string {
	// Implement me
	var duplicate []string
	individual_chars := strings.Split(original, "")
	twice := 2
	for _, individual_char := range individual_chars {
		i := 0
		for i < twice {
			duplicate = append(duplicate, individual_char)
			i += 1
		}
	}
	return strings.Join(duplicate, "")
}

func main() {
	fmt.Println("Smallest Difference Challenge")

	original := "gophers"
	doubled := DoubleChars(original)
	fmt.Println(doubled) // ggoopphheerrss
}
