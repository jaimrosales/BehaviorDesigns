package main

import "fmt"

func main() {
	listaNumber := &Number{}
	listaWord := &Word{}

	listaNumber.Add(1)
	listaNumber.Add(2)
	listaNumber.Add(4)
	listaNumber.Add(8)

	listaWord.Set(0, "Hola")
	listaWord.Set(1, "como")
	listaWord.Set(2, "estas")
	listaWord.Set(3, "?")

	//numberIt := listaNumber.GetIterator()
	//wordIt := listaWord.GetIterator()

	var iterador MyIterator

	iterador = listaNumber.GetIterator()

	iterador.First()
	for iterador.HasMore() {
		fmt.Println(iterador.Next())
	}
	println("-----")
	iterador = listaWord.GetIterator()
	iterador.First()
	for iterador.HasMore() {
		println(iterador.Next().(string))
	}

}
