package main

//Son ejemplos de estructuras complejas,no tiene importancia al patron de diseno
//Esta es una estructura compleja que tiene un slice de numeros
type Number struct {
	Numbers []int
}

//Metodo para agregar numeros a la estructura
func (n *Number) Add(m int) {
	n.Numbers = append(n.Numbers, m)
}

func (n *Number) GetIterator() *numberIterator {
	return &numberIterator{theNumber: n}
}
