package more

var p *person

func GetInstance() *person {
	if p == nil {
		p = &person{}
	}
	return p
}

type person struct {
	name string
	age  int
}
