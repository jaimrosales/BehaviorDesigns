package main

type Person struct {
	Name      string
	Observers []Observer
}

func (p *Person) SetName(name string) {
	p.Name = name
	p.NotifyObservers()
}

func (p *Person) AddObserver(o Observer) {
	// Implementation to add an observer
	p.Observers = append(p.Observers, o)
}

func (p *Person) NotifyObservers() {
	for _, observer := range p.Observers {
		observer.notify(p.Name)
	}
}
