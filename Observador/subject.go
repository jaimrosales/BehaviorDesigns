package main

type Subject interface {
	AddObserver(o Observer)
	//DeleteObserver(o Observer)
	NotifyObservers()
}
