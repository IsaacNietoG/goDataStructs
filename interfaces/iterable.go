package interfaces

import "iter"

type Iterable[V any] interface{
	All() iter.Seq[V]
}
