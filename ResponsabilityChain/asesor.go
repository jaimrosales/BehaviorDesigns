package main

type Asesor struct {
	sucesor Autoriza
}

func (a *Asesor) SetSucesor(s Autoriza) {
	a.sucesor = s
}

func (a *Asesor) Autorizar(monto int) {
	if monto <= 2000 {
		//Logica de autorizacion
		println("Asesor autoriza el monto de: ", monto)
	} else {
		println("Asesor no puede autorizar el monto de: ", monto)
		if a.sucesor != nil {
			a.sucesor.Autorizar(monto)
		} else {
			println("No se puede autorizar el monto de: ", monto)
		}
	}
}
