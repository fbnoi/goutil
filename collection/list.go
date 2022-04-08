package collection

type IList interface {
	Collection

	AddAt(p int, e interface{})
	Get(p int) interface{}
	RemoveAt(p int)
	Set(p int, e interface{})
}
