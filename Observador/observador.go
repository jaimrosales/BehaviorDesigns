package main

type Observer interface {
	notify(event string)
}
