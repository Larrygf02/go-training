package client_two

import "github.com/larrygf02/patterns/singlenton"

func IncrementAge() {
	p := singlenton.GetInstance()
	p.IncrementAge()
}
