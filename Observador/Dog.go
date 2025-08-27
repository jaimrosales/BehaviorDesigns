package main

import "fmt"

type Dog struct {
	age       int
	Observers []Observer
}

func (d *Dog) SetAge(age int) {
	d.age = age
	d.NotifyObservers()
}

func (d *Dog) AddObserver(o Observer) {
	// Implementation to add an observer
	d.Observers = append(d.Observers, o)
}

func (d *Dog) NotifyObservers() {
	for _, observer := range d.Observers {
		observer.notify(fmt.Sprintf("Dog's age changed to %v", d.age))
	}
}
