package interfaces


import "iter"

type Coleccion[V any] interface{
	Agrega(elemento V) (err Error)
	Elimina(elemento V) (err Error)
	Contiene(elemento V) (b bool, err Error)
	EsVacia() (b bool, err Error)
	GetElementos() (i int, err Error)
	Limpia() (err Error)
	All() iter.Seq[V]
}
