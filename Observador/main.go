package main

func main() {
	p := &Person{}
	s := &Slack{}
	m := &Email{}

	p.AddObserver(m)
	p.AddObserver(s)
	p.SetName("Jaime")

	d := &Dog{}
	d.AddObserver(m)

	d.SetAge(5)

}
