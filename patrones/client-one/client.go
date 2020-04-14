package client_one

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
