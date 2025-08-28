package main

type Coordinador struct {
	sucesor Autoriza
}

func (c *Coordinador) SetSucesor(s Autoriza) {
	c.sucesor = s
}

func (c *Coordinador) Autorizar(monto int) {
	if monto <= 5000 {
		//Logica de autorizacion
		println("Coordinador autoriza el monto de: ", monto)
	} else {
		println("Coordinador no puede autorizar el monto de: ", monto)
		if c.sucesor != nil {
			c.sucesor.Autorizar(monto)
		} else {
			println("No se puede autorizar el monto de: ", monto)
		}
	}
}
