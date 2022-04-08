package collection

type IList[T comparable] interface {
	Collection[T]

	AddAt(p int, e T)
	Get(p int) T
	RemoveAt(p int)
	Set(p int, e T)
}
