package main

import "fmt"

type Stack []Flight

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (s *Stack) Pop() Flight {
	// TODO Implement
	if s.IsEmpty() {
		return Flight{}
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func (s *Stack) Push(f Flight) {
	// TODO Implement
	*s = append(*s, f)
}

func (s *Stack) Peek() Flight {
	// TODO Implement
	return Flight{}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func main() {
	fmt.Println("Go Stack Implementation")
	var stack Stack
	stack.Push(Flight{Origin: "a"})
	stack.Push(Flight{Origin: "b"})
	stack.Push(Flight{Origin: "c"})
	stack.Push(Flight{Origin: "d"})
	for len(stack) > 0 {
		x := stack.Pop()
		fmt.Println(x)
	}
}
