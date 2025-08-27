package main

import (
	"fmt"
)

type Email struct{}

func (e *Email) notify(event string) {
	fmt.Println("Email notification:", event)
}
