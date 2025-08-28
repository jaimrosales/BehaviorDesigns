package main

import "fmt"

func main() {
	//El ejemplo del patron cadena de responsabilidad permite pasar una solicitud a diferentes elementos hasta que uno de ellos pueda procesarla
	//elementos de la cadena
	a := &Asesor{}
	c := &Coordinador{}
	g := &gerente{}
	//configuracion de la cadena
	a.SetSucesor(c)
	c.SetSucesor(g)
	//solicitudes
	d := 0
	fmt.Println("Solicitud de prestamo por: ")
	fmt.Scanln(&d)
	a.Autorizar(d)
}
