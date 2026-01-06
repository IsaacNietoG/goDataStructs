package lists
import (
	"iter"
	"github.com/IsaacNietoG/goDataStructs/interfaces/Coleccion"
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

func (l *SinglyLinkedList[V]) Agrega(elemento V) (error){
	
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

func (l *SinglyLinkedList[V]) Elimina(elemento V) (error){
	
}

func (l *SinglyLinkedList[V]) EliminaFunc(elemento V, cmp func(V,V)) (error){
	
}

func (l *SinglyLinkedList[V]) Contiene(elemento V) (bool,error) {
	
}

func (l *SinglyLinkedList[V]) ContieneFunc(elemento V, cmp func(V,V)) (bool,error) {
	
}

func (l *SinglyLinkedList[V]) EsVacia() (bool, error) {
	
}

func (l *SinglyLinkedList[V]) GetElementos() (int, error) {
	
}

func (l *SinglyLinkedList[V]) Limpia() (error){
	
}

func (l *SinglyLinkedList[V]) All() iter.Seq[V] {
	
}

func (l *SinglyLinkedList[V]) Equal(coleccion Coleccion[V]) (bool,error){
	
}

func (l *SinglyLinkedList[V]) EqualFunc(coleccion Coleccion[V], cmp func(V,V)) (bool, error){
	
}

func (l *SinglyLinkedList[V]) String() (string){
	
}

func (l *SinglyLinkedList[V]) AgregaInicio() (error){
	
}

func (l *SinglyLinkedList[V]) Inserta(indice int, elemento V) (error){
	
}

func (l *SinglyLinkedList[V]) EliminaPrimero() (V, error){
	
}

func (l *SinglyLinkedList[V]) EliminaUltimo() (V, error){
	
}

func (l *SinglyLinkedList[V]) GetPrimero() (V, error){
	
}

func (l *SinglyLinkedList[V]) GetUltimo() (V, error){
	
}

func (l *SinglyLinkedList[V]) Copia() (*SinglyLinkedList[V], error){
	
}

func (l *SinglyLinkedList[V]) Get(i int) (V, error){
	
}

func (l *SinglyLinkedList[V]) IndiceDe(elemento V) (int, error){
	
}

func (l *SinglyLinkedList[V]) IndiceDeFunc(elemento V, cmp func(V,V)) (int, error){
	
}
