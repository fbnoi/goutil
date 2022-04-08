package collection

type Collection interface {
	Iterable

	Add(e interface{}) bool
	AddAll(e Collection) bool
	Clear()
	Contains(e interface{}) bool
	ContainsAll(e Collection) bool
	Empty() bool
	Remove(e interface{}) bool
	RemoveAll(e Collection) bool
	RetainAll(e Collection) bool
	Size() int
	ToSlice() []interface{}
}

type Iterable interface {
	GetIterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
	Remove()
}
