package lists

import (
	"iter"
	"github.com/IsaacNietoG/goDataStructs/interfaces/coleccion"
)

type nodoSimple[V any] struct {
	elemento V
	siguiente *nodoSimple[V]
}

type SinglyLinkedList[V any] struct{
	head *nodoSimple[V]
	tail *nodoSimple[V]
	size int
}

var _ Coleccion[any] = (*SinglyLinkedList[any])(nil)

func (l *SinglyLinkedList[V]) Agrega(elemento V) (err Error){
	
	n:= &nodoSimple[V]{elemento: elemento}
	l.size++
	
	if l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.siguiente = n
		l.tail = n
	}
	
	return nil
	
}

func (l *SinglyLinkedList[V]) Elimina(elemento V) (err Error){
	
}

func (l *SinglyLinkedList[V]) Contiene(elemento V) (b bool, err Error) {
	
}

func (l *SinglyLinkedList[V]) EsVacia() (b bool, err Error) {
	
}

func (l *SinglyLinkedList[V]) GetElementos() (i int, err Error) {
	
}

func (l *SinglyLinkedList[V]) Limpia() (err Error){
	
}

func (l *SinglyLinkedList[V]) All() iter.Seq[V] {
	
}

func (l *SinglyLinkedList[V]) AgregaInicio() (err Error){
	
}

func (l *SinglyLinkedList[V]) Inserta(indice int, elemento V) (err Error){
	
}

func (l *SinglyLinkedList[V]) EliminaPrimero() (elemento V, err Error){
	
}

func (l *SinglyLinkedList[V]) EliminaUltimo() (elemento V, err Error){
	
}

func (l *SinglyLinkedList[V]) GetPrimero() (elemento V, err Error){
	
}

func (l *SinglyLinkedList[V]) GetUltimo() (elemento V, err Error){
	
}

func (l *SinglyLinkedList[V]) Copia() (copia *SinglyLinkedList[V], err Error){
	
}

func (l *SinglyLinkedList[V]) Get(i int) (elemento V, err Error){
	
}

func (l *SinglyLinkedList[V]) IndiceDe(elemento *V) (i int, err Error){
	
}

func Equal[T comparable](l1, l2 *SinglyLinkedList[T]) (b bool, err Error){
	
}
func EqualFunc[E1, E2 any](l1 *SinglyLinkedList[E1], l2 *SinglyLinkedList[E2], eq func (E1, E2) bool) (b bool, err Error){
	
}

func (l *SinglyLinkedList[V]) String() (s String, err Error){
	
}
