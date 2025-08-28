package main

type MyIterator interface {
	//Todas las estructuras que se quieran iterar tienen que implementar estos metodos
	First()            //Va al primer elemento de la estructura en la que se esta iterando
	Next() interface{} //Va al siguiente elemento de la estructura en la que se esta iterando, tiene que devolver una interface vacia devido a que no se sabe que tipo de dato va a devolver
	HasMore() bool     //Devuelve true si hay mas elementos en la estructura en la que se esta iterando
	
}
