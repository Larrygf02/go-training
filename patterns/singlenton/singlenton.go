package singlenton

import "sync"

var p *person
var once sync.Once

func GetInstance() *person {
	// se ejecuta una sola vez "Singlenton"
	once.Do(func() {
		p = &person{}
	})
	return p
}

type person struct {
	name string
	age  int
	mux  sync.Mutex
}

func (p *person) SetName(name string) {
	p.mux.Lock()
	p.name = name
	p.mux.Unlock()
}

func (p *person) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *person) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *person) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age

}
