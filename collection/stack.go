package collection

type Stack struct {
	LinkedList
}

func (stack *Stack) Push(val interface{}) interface{} {
	stack.linkFirst(val)
	return val
}

func (stack *Stack) Pop() interface{} {
	if stack.Empty() {
		panic("no such element")
	}
	val := stack.Get(0)
	stack.unlinkFirst()
	return val
}

func (stack *Stack) Top() interface{} {
	if stack.Empty() {
		panic("no such element")
	}
	val := stack.Get(0)
	return val
}

func (stack *Stack) Search(val interface{}) int {
	return stack.IndexOf(val)
}
