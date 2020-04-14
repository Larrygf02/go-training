package singlenton

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

func (p *person) SetName(name string) {
	p.name = name
}

func (p *person) GetName() string {
	return p.name
}

func (p *person) IncrementAge() {
	p.age++
}

func (p *person) GetAge() int {
	return p.age
}
