package main

import "fmt"

type Slack struct {
}

func (s *Slack) notify(event string) {
	fmt.Println("Slack notification:", event)
}
