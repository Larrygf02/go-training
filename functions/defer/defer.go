package main
import (
	"fmt"
)

func main() {
	/* Las llamadas a funciones diferidas se envían a una
		una pila (ULTIMO EN ENTRAR PRIMERO EN SALIR)
		Las llamadas diferidas no se ejecutan hasta que se
		devuelva la función circundante
	*/
	defer foo()
	defer bar()
	hello("Grace")
}

func foo() {
	fmt.Println("Foo")
}

func bar() {
	fmt.Println("Bar")
}

func hello(name string) {
	fmt.Println("Hello", name)
}