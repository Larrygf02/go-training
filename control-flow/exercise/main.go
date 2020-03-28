package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
)

type usercoin struct {
	user   string
	coints int
}

func main() {
	for _, user := range users {
		coin := 0
		for _, char := range user {
			switch char {
			case 'a', 'e':
				coin += 1
			case 'i':
				coin += 2
			case 'o':
				coin += 3
			case 'u':
				coin += 4
			}
		}
		if coin > 10 {
			coin = 10
		}
		newusercoin := usercoin{
			user,
			coin,
		}
		fmt.Println(newusercoin)
		// break
	}
}
