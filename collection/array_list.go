package collection

type ArrayList[T comparable] struct {
	array    []T
	modCount int
}

func (al *ArrayList[T]) Add(e T) bool {
	al.array = append(al.array, e)
	al.modCount++
	return true
}

func (al *ArrayList[T]) AddAt(p int, e T) {
	if p < 0 || p > al.Size()-1 {
		panic("list bounds out of range")
	}
	var np []T
	for i := 0; i < al.Size(); i++ {
		if i == p {
			np = append(np, e)
		}
		np = append(np, al.Get(i))
	}
	al.array = np
	al.modCount++
}

func (al *ArrayList[T]) AddAll(c Collection[T]) bool {
	if c.Size() <= 0 {
		return false
	}
	slic := c.ToSlice()
	al.array = append(al.array, slic...)
	al.modCount = al.modCount + len(slic)
	return true
}

func (al *ArrayList[T]) Clear() {
	al.array = []T{}
	al.modCount = 0
}

func (al *ArrayList[T]) Contains(e T) bool {
	return al.IndexOf(e) > -1
}

func (al *ArrayList[T]) ContainsAll(c Collection[T]) bool {
	it := c.GetIterator()
	for it.HasNext() {
		if !al.Contains(it.Next()) {
			return false
		}
	}
	return true
}

func (al *ArrayList[T]) Get(p int) T {
	if p < 0 || p > al.Size()-1 {
		panic("list bounds out of range")
	}
	return al.array[p]
}

func (al *ArrayList[T]) RemoveAt(p int) {
	al.Remove(al.Get(p))
}

func (al *ArrayList[T]) Set(p int, e T) {
	if p < 0 || p > al.Size()-1 {
		panic("list bounds out of range")
	}
	al.array[p] = e
}

func (al *ArrayList[T]) Empty() bool {
	return len(al.array) == 0
}

func (al *ArrayList[T]) Remove(e T) bool {
	idx := al.IndexOf(e)
	if idx == -1 {
		return false
	}
	if idx == len(al.array)-1 {
		al.array = al.array[:idx]
		al.modCount++
	} else if idx != -1 {
		al.array = append(al.array[:idx], al.array[idx+1:]...)
		al.modCount++
	}
	return idx != -1
}

func (al *ArrayList[T]) RemoveAll(c Collection[T]) bool {
	var np []T
	for _, el := range al.array {
		if !c.Contains(el) {
			np = append(np, el)
		}
	}
	s := len(al.array)
	w := len(np)
	if s != w {

		al.array = np
		al.modCount = al.modCount + s - w

		return true
	}
	return false
}

func (al *ArrayList[T]) RetainAll(c Collection[T]) bool {
	var np []T
	for _, el := range al.array {
		if c.Contains(el) {
			np = append(np, el)
		}
	}
	s := len(al.array)
	w := len(np)
	if s != w {

		al.array = np
		al.modCount = al.modCount + s - w

		return true
	}
	return false
}

func (al *ArrayList[T]) Size() int {
	return len(al.array)
}

func (al *ArrayList[T]) ToSlice() []T {
	dest := make([]T, al.Size())
	copy(dest, al.array)
	return dest
}

func (al *ArrayList[T]) GetIterator() Iterator[T] {
	return &AlIterator[T]{
		idx:            -1,
		al:             al,
		expectModCount: al.modCount,
	}
}

func (al *ArrayList[T]) IndexOf(e T) int {
	for i, el := range al.array {
		if el == e {
			return i
		}
	}
	return -1
}

type AlIterator[T comparable] struct {
	idx            int
	al             *ArrayList[T]
	expectModCount int
}

func (it *AlIterator[T]) HasNext() bool {
	it.checkStatus()
	return it.idx < len(it.al.array)-1
}

func (it *AlIterator[T]) Next() T {
	it.checkStatus()
	it.idx++
	return it.al.Get(it.idx)
}

func (it *AlIterator[T]) Remove() {
	it.al.Remove(it.al.Get(it.idx))
	it.expectModCount++
	it.idx--
}

func (it *AlIterator[T]) checkStatus() {
	if it.expectModCount != it.al.modCount {
		panic("list has been modified during Iterate")
	}
}
