package interfaces


type Coleccion[V any] interface{
	agrega(elemento V)
	elimina(elemento V)
	contiene(elemento V) bool
}
