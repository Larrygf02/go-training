package main
import (
	"encoding/json"
	"fmt"
)

type message struct {
	Sender string
	Body string
	To string
}

type jsonMessage struct {
	Sender string `json:"Sender"`
	Body   string `json:"Body"`
	To     string `json:"To"`
}
func main() {
	newmessage := message {
		Sender: "Raul",
		Body: "Hello",
		To: "Grace",
	}

	othermessage := message {
		Sender: "Grace",
		Body: "Hi",
		To: "Raul",
	}

	messages := [] message{newmessage, othermessage}
	fmt.Println(messages)
	jm, _ := json.Marshal(messages)
	fmt.Println(string(jm))

	// Unmarshal
	s := `[{"Sender":"Raul","Body":"Hello","To":"Grace"},{"Sender":"Grace","Body":"Hi","To":"Raul"}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var ms []jsonMessage

	err := json.Unmarshal(bs, &ms)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ms)
}