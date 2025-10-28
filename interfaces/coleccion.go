package interfaces


type Coleccion[V any] interface{
	agrega(elemento V) void
	elimina(elemento V) void
	contiene(elemento V) bool
}
