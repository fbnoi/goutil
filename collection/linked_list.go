package collection

type entry[T comparable] struct {
	previous *entry[T]
	next     *entry[T]
	value    T
}

type LinkedList[T comparable] struct {
	head     *entry[T]
	end      *entry[T]
	size     int
	modCount int
}

func (ll *LinkedList[T]) linkFirst(val T) {
	entry := &entry[T]{
		next:  ll.head,
		value: val,
	}
	if ll.head != nil {
		ll.head.previous = entry
	} else {
		ll.end = entry
	}
	ll.head = entry
	ll.size++
	ll.modCount++
}

func (ll *LinkedList[T]) linkLast(val T) {
	entry := &entry[T]{
		previous: ll.end,
		value:    val,
	}
	if ll.end != nil {
		ll.end.next = entry
	} else {
		ll.head = entry
	}
	ll.end = entry
	ll.size++
	ll.modCount++
}

func (ll *LinkedList[T]) linkbefore(val T, e *entry[T]) {
	entry := &entry[T]{
		value:    val,
		next:     e,
		previous: e.previous,
	}
	if e.previous != nil {
		e.previous.next = entry
	}
	e.previous = entry
	if e == ll.head {
		ll.head = entry
	}
	ll.size++
	ll.modCount++
}

func (ll *LinkedList[T]) unlinkFirst() bool {
	if ll.head != nil {
		next := ll.head.next
		head := ll.head
		if head == ll.end {
			ll.end = nil
		}
		if next != nil {
			next.previous = nil
		}
		head.value = zero[T]()
		head.next = nil
		ll.head = next
		ll.size--
		ll.modCount++
		return true
	}
	return false
}

func (ll *LinkedList[T]) unlinkLast() bool {
	if ll.end != nil {
		previous := ll.end.previous
		end := ll.end
		if end == ll.head {
			ll.head = nil
		}
		if previous != nil {
			previous.next = nil
		}
		end.value = zero[T]()
		end.previous = nil
		ll.end = previous
		ll.size--
		ll.modCount++
		return true
	}
	return false
}

func (ll *LinkedList[T]) unlink(en *entry[T]) bool {
	if en == nil {
		return false
	}
	if en == ll.end {
		ll.unlinkLast()
	} else if en == ll.head {
		ll.unlinkFirst()
	} else {
		en.previous.next = en.next
		en.next.previous = en.previous
		en.next = nil
		en.previous = nil
		en.value = zero[T]()
		ll.size--
		ll.modCount++
	}
	return true
}

func (ll *LinkedList[T]) checkIndex(p int) {
	if p < 0 || p > ll.Size()-1 {
		panic("list bounds out of range")
	}
}

func (ll *LinkedList[T]) Add(e T) bool {
	ll.linkLast(e)
	return true
}

func (ll *LinkedList[T]) AddFirst(e T) {
	ll.linkFirst(e)
}

func (ll *LinkedList[T]) AddAt(p int, e T) {
	ll.checkIndex(p)
	en := ll.head
	for j := 0; en != nil && j < p; j++ {
		en = en.next
	}
	if en == nil {
		ll.linkLast(e)
	} else {
		ll.linkbefore(e, en)
	}
}

func (ll *LinkedList[T]) AddAll(c Collection[T]) bool {
	if c.Size() > 0 {
		it := c.GetIterator()
		for it.HasNext() {
			ll.linkLast(it.Next())
		}
		return true
	}
	return false
}

func (ll *LinkedList[T]) AddAllAt(p int, c Collection[T]) bool {
	ll.checkIndex(p)
	if c.Size() > 0 {
		en := ll.head
		for j := 0; en != nil && j < p; j++ {
			en = en.next
		}
		it := c.GetIterator()
		for it.HasNext() {
			if en == nil {
				ll.linkLast(it.Next())
			} else {
				ll.linkbefore(it.Next(), en)
			}
		}
	}
	return false
}

func (ll *LinkedList[T]) Clear() {
	e := ll.head
	for e != nil {
		next := e.next
		e.next = nil
		e.previous = nil
		e.value = zero[T]()
		e = next
	}
	ll.head = nil
	ll.end = nil
	ll.size = 0
	ll.modCount++
}

func (ll *LinkedList[T]) Reset() {
	e := ll.head
	for e != nil {
		next := e.next
		e.next = nil
		e.previous = nil
		e.value = zero[T]()
		e = next
	}
	ll.head = nil
	ll.end = nil
	ll.size = 0
	ll.modCount = 0
}

func (ll *LinkedList[T]) IndexOf(e T) int {
	for i, en := 0, ll.head; i < ll.size && en != nil; i, en = i+1, en.next {
		if en.value == e {
			return i
		}
	}
	return -1
}

func (ll *LinkedList[T]) Contains(e T) bool {
	return ll.IndexOf(e) > -1
}

func (ll *LinkedList[T]) ContainsAll(c Collection[T]) bool {
	it := c.GetIterator()
	findAll := true
	for it.HasNext() {
		if ll.IndexOf(it.Next()) < 0 {
			findAll = false
		}
	}
	return findAll
}

func (ll *LinkedList[T]) Get(p int) T {
	ll.checkIndex(p)
	en := ll.head
	for j := 0; en != nil && j < p; j, en = j+1, en.next {
	}
	return en.value
}

func (ll *LinkedList[T]) Empty() bool {
	return ll.size == 0
}

func (ll *LinkedList[T]) Remove(e T) bool {
	en := ll.head
	var f *entry[T]
	for ; en != nil; en = en.next {
		if en.value == e {
			f = en
			break
		}
	}
	return ll.unlink(f)
}

func (ll *LinkedList[T]) RemoveAt(p int) {
	ll.checkIndex(p)
	en := ll.head
	for j := 0; en != nil && j < p; j, en = j+1, en.next {
	}
	ll.unlink(en)
}

func (ll *LinkedList[T]) RemoveAll(c Collection[T]) bool {
	modify := false
	entry := ll.head
	for entry != nil {
		next := entry.next
		if c.Contains(entry.value) {
			ll.unlink(entry)
			modify = true
		}
		entry = next
	}
	return modify
}

func (ll *LinkedList[T]) Set(p int, e T) {
	ll.checkIndex(p)
	en := ll.head
	for j := 0; en != nil && j < p; j, en = j+1, en.next {
	}
	en.value = e
	ll.modCount++
}

func (ll *LinkedList[T]) RetainAll(c Collection[T]) bool {
	entry := ll.head
	modify := false
	for entry != nil {
		next := entry.next
		if !c.Contains(entry.value) {
			ll.unlink(entry)
			modify = true
		}
		entry = next
	}
	return modify
}

func (ll *LinkedList[T]) Size() int {
	return ll.size
}

func (ll *LinkedList[T]) ToSlice() []T {
	var dest []T
	entry := ll.head
	for entry != nil {
		dest = append(dest, entry.value)
		entry = entry.next
	}
	return dest
}

func (ll *LinkedList[T]) GetIterator() Iterator[T] {
	return &lIterator[T]{
		next:             ll.head,
		expectedModCount: ll.modCount,
		list:             ll,
		index:            -1,
		lastReturn:       nil,
	}
}

type lIterator[T comparable] struct {
	index            int
	lastReturn       *entry[T]
	next             *entry[T]
	expectedModCount int
	list             *LinkedList[T]
}

func (li *lIterator[T]) HasNext() bool {
	return li.index < li.list.size-1
}

func (li *lIterator[T]) Next() T {
	li.checkModStatus()
	if !li.HasNext() {
		panic("no such elem")
	}
	li.lastReturn = li.next
	li.next = li.next.next
	li.index++
	return li.lastReturn.value
}

func (li *lIterator[T]) Remove() {
	if li.lastReturn == nil {
		li.invalidLastReturn()
	}
	lastNext := li.lastReturn.next
	li.list.unlink(li.lastReturn)
	if li.next == li.lastReturn {
		li.next = lastNext
	} else {
		li.index--
	}
	li.lastReturn = nil
	li.expectedModCount++
}

func (li *lIterator[T]) Set(e T) {
	if li.lastReturn == nil {
		li.invalidLastReturn()
	}
	li.checkModStatus()
	li.lastReturn.value = e
}

func (li *lIterator[T]) Add(e T) {
	if li.lastReturn == nil {
		li.invalidLastReturn()
	}
	li.checkModStatus()
	if li.next == nil {
		li.list.linkLast(e)
	} else {
		li.list.linkbefore(e, li.next)
	}
	li.lastReturn.value = e
}

func (li *lIterator[T]) checkModStatus() {
	if li.expectedModCount != li.list.modCount {
		panic("list has been modified during Iterate")
	}
}

func (li *lIterator[T]) invalidLastReturn() {
	panic("last returns is nil")
}
