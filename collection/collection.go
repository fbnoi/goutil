package collection

type Collection[T comparable] interface {
	Iterable[T]

	Add(e T) bool
	AddAll(e Collection[T]) bool
	Clear()
	Contains(e T) bool
	ContainsAll(e Collection[T]) bool
	Empty() bool
	Remove(e T) bool
	RemoveAll(e Collection[T]) bool
	RetainAll(e Collection[T]) bool
	Size() int
	ToSlice() []T
}

type Iterable[T comparable] interface {
	GetIterator() Iterator[T]
}

type Iterator[T comparable] interface {
	HasNext() bool
	Next() T
	Remove()
}

func zero[T any]() T {
	var zero T
	return zero
}
