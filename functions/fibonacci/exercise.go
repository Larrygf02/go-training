package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	return func(x int) int{
		if x == 0 {
			return 0
		}else if x == 1 {
			return 1
		}else {
			f := fibonacci()
			return f(x-1) + f(x-2);
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}