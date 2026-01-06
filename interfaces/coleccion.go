package interfaces

import "iter"

type Coleccion[V any] interface{
	Agrega(elemento V) (error)
	Elimina(elemento V) (error)
	EliminaFunc(elemento V, cmp func(V,V)) (error)
	Contiene(elemento V) (bool, error)
	ContieneFunc(elemento V, cmp func(V,V)) (bool, error)
	EsVacia() (bool, error)
	GetElementos() (int, error)
	Limpia() (error)
	All() iter.Seq[V]
	Equal(coleccion Coleccion[V]) (bool, error)
	EqualFunc(coleccion Coleccion[V], cmp func(V,V)) (bool, error)
	String() (string)
}
