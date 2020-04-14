package client_one

import "github.com/larrygf02/patterns/singlenton"

func IncrementAge() {
	p := singlenton.GetInstance()
	p.IncrementAge()
}
