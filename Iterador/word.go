package main

//Son ejemplos de estructuras complejas,no tiene importancia al patron de diseno
//Esta es una estructura compleja que tiene 4 palabras
type Word struct {
	Word1 string
	Word2 string
	Word3 string
	Word4 string
}

func (wi *Word) GetIterator() *wordIterator {
	return &wordIterator{theWord: wi}

}

//Metodo para agregar palabras a la estructura
func (w *Word) Set(i int, s string) {
	switch i {
	case 0:
		w.Word1 = s
	case 1:
		w.Word2 = s
	case 2:
		w.Word3 = s
	case 3:
		w.Word4 = s
	}
}
