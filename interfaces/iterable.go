package interfaces


import "golang/x/exp/iter"

type Iterable[V any] interface{
	All() iter.Seq[V]
}
