package collection

type ArrayList struct {
	array    []interface{}
	modCount int
}

func (al *ArrayList) Add(e interface{}) bool {
	al.array = append(al.array, e)
	al.modCount++
	return true
}

func (al *ArrayList) AddAt(p int, e interface{}) {
	if p < 0 || p > al.Size()-1 {
		panic("list bounds out of range")
	}
	var np []interface{}
	for i := 0; i < al.Size(); i++ {
		if i == p {
			np = append(np, e)
		}
		np = append(np, al.Get(i))
	}
	al.array = np
	al.modCount++
}

func (al *ArrayList) AddAll(c Collection) bool {
	if c.Size() <= 0 {
		return false
	}
	slic := c.ToSlice()
	al.array = append(al.array, slic...)
	al.modCount = al.modCount + len(slic)
	return true
}

func (al *ArrayList) Clear() {
	al.array = []interface{}{}
	al.modCount = 0
}

func (al *ArrayList) Contains(e interface{}) bool {
	return al.IndexOf(e) > -1
}

func (al *ArrayList) ContainsAll(c Collection) bool {
	it := c.GetIterator()
	for it.HasNext() {
		if !al.Contains(it.Next()) {
			return false
		}
	}
	return true
}

func (al *ArrayList) Get(p int) interface{} {
	if p < 0 || p > al.Size()-1 {
		panic("list bounds out of range")
	}
	return al.array[p]
}

func (al *ArrayList) RemoveAt(p int) {
	al.Remove(al.Get(p))
}

func (al *ArrayList) Set(p int, e interface{}) {
	if p < 0 || p > al.Size()-1 {
		panic("list bounds out of range")
	}
	al.array[p] = e
}

func (al *ArrayList) Empty() bool {
	return len(al.array) == 0
}

func (al *ArrayList) Remove(e interface{}) bool {
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

func (al *ArrayList) RemoveAll(c Collection) bool {
	var np []interface{}
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

func (al *ArrayList) RetainAll(c Collection) bool {
	var np []interface{}
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

func (al *ArrayList) Size() int {
	return len(al.array)
}

func (al *ArrayList) ToSlice() []interface{} {
	dest := make([]interface{}, al.Size())
	copy(dest, al.array)
	return dest
}

func (al *ArrayList) GetIterator() Iterator {
	return &AlIterator{
		idx:            -1,
		al:             al,
		expectModCount: al.modCount,
	}
}

func (al *ArrayList) IndexOf(e interface{}) int {
	for i, el := range al.array {
		if el == e {
			return i
		}
	}
	return -1
}

type AlIterator struct {
	idx            int
	al             *ArrayList
	expectModCount int
}

func (it *AlIterator) HasNext() bool {
	it.checkStatus()
	return it.idx < len(it.al.array)-1
}

func (it *AlIterator) Next() interface{} {
	it.checkStatus()
	it.idx++
	return it.al.Get(it.idx)
}

func (it *AlIterator) Remove() {
	it.al.Remove(it.al.Get(it.idx))
	it.expectModCount++
	it.idx--
}

func (it *AlIterator) checkStatus() {
	if it.expectModCount != it.al.modCount {
		panic("list has been modified during Iterate")
	}
}
