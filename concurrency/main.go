package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Gorutines\t", runtime.NumGoroutine())
	go foo()
	fmt.Println("Gorutines\t", runtime.NumGoroutine())
}

func foo() {
	fmt.Println("Go")
}
