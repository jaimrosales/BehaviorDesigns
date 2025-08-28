package main

type gerente struct {
	sucesor Autoriza
}

func (g *gerente) SetSucesor(s Autoriza) {
	g.sucesor = s
}
func (g *gerente) Autorizar(monto int) {
	if monto <= 10000 {
		//Logica de autorizacion
		println("Gerente autoriza el monto de: ", monto)
	} else {
		println("Gerente no puede autorizar el monto de: ", monto)
		if g.sucesor != nil {
			g.sucesor.Autorizar(monto)
		} else {
			println("No se puede autorizar el monto de: ", monto)
		}
	}
}
