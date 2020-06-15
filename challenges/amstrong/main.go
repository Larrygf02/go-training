package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type MyInt int

// IsArmstrong checks to see if a number is
// indeed an armstrong number and returns a
// bool
func (n *MyInt) IsArmstrong() bool {
	number_string := strconv.Itoa(int(*n))
	digits_numbers := strings.Split(number_string, "")
	sum := 0
	for _, digit_string := range digits_numbers {
		digit, _ := strconv.Atoi(digit_string)
		sum += int(math.Pow(float64(digit), 3))
	}
	fmt.Println(sum)
	if sum == int(*n) {
		return true
	}
	return false
}

func main() {
	fmt.Println("Armstrong Numbers")

	var num1 MyInt = 371
	fmt.Println(num1.IsArmstrong())
}
