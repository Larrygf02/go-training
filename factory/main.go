package main

import (
	"fmt"
	"os"

	"github.com/larrygf02/factory/factory"
)

func main() {
	var t int
	fmt.Print("Digite la conexión desea 1: MYSQL, 2: POSTGRES ==> ")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("error al leer la opcion")
		os.Exit(1)
	}

	conn := factory.Factory(t)
	err = conn.Connect()
	if err != nil {
		fmt.Printf("No se pudo crear la conexion %v", err)
		os.Exit(1)
	}
	if conn == nil {
		fmt.Println("Motor no valido")
		os.Exit(1)
	}

	now, err := conn.GetNow()
	if err != nil {
		fmt.Printf("No se pudo consultar la fecha %v", err)
	}
	fmt.Println(now.Format("2006-01-02"))
	err = conn.Close()
	if err != nil {
		fmt.Printf("No se pudo cerrar la conexion %v", err)
	}

}
