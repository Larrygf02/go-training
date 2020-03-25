package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	fmt.Println(s)
	fmt.Println(bs)

	login := `password123`
	err := bcrypt.CompareHashAndPassword(bs, []byte(login))
	if err != nil {
		fmt.Println("No te puedes logear")
	}

	fmt.Println("Te puedes loguear")
}