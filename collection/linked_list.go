package collection

type entry struct {
	previous *entry
	next     *entry
	value    interface{}
}

type LinkedList struct {
	head     *entry
	end      *entry
	size     int
	modCount int
}

func (ll *LinkedList) linkFirst(val interface{}) {
	entry := &entry{
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

func (ll *LinkedList) linkLast(val interface{}) {
	entry := &entry{
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

func (ll *LinkedList) linkbefore(val interface{}, e *entry) {
	entry := &entry{
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

func (ll *LinkedList) unlinkFirst() bool {
	if ll.head != nil {
		next := ll.head.next
		head := ll.head
		if head == ll.end {
			ll.end = nil
		}
		if next != nil {
			next.previous = nil
		}
		head.value = nil
		head.next = nil
		ll.head = next
		ll.size--
		ll.modCount++
		return true
	}
	return false
}

func (ll *LinkedList) unlinkLast() bool {
	if ll.end != nil {
		previous := ll.end.previous
		end := ll.end
		if end == ll.head {
			ll.head = nil
		}
		if previous != nil {
			previous.next = nil
		}
		end.value = nil
		end.previous = nil
		ll.end = previous
		ll.size--
		ll.modCount++
		return true
	}
	return false
}

func (ll *LinkedList) unlink(en *entry) bool {
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
		en.value = nil
		ll.size--
		ll.modCount++
	}
	return true
}

func (ll *LinkedList) checkIndex(p int) {
	if p < 0 || p > ll.Size()-1 {
		panic("list bounds out of range")
	}
}

func (ll *LinkedList) Add(e interface{}) bool {
	ll.linkLast(e)
	return true
}

func (ll *LinkedList) AddFirst(e interface{}) {
	ll.linkFirst(e)
}

func (ll *LinkedList) AddAt(p int, e interface{}) {
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

func (ll *LinkedList) AddAll(c Collection) bool {
	if c.Size() > 0 {
		it := c.GetIterator()
		for it.HasNext() {
			ll.linkLast(it.Next())
		}
		return true
	}
	return false
}

func (ll *LinkedList) AddAllAt(p int, c Collection) bool {
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

func (ll *LinkedList) Clear() {
	e := ll.head
	for e != nil {
		next := e.next
		e.next = nil
		e.previous = nil
		e.value = nil
		e = next
	}
	ll.head = nil
	ll.end = nil
	ll.size = 0
	ll.modCount++
}

func (ll *LinkedList) Reset() {
	e := ll.head
	for e != nil {
		next := e.next
		e.next = nil
		e.previous = nil
		e.value = nil
		e = next
	}
	ll.head = nil
	ll.end = nil
	ll.size = 0
	ll.modCount = 0
}

func (ll *LinkedList) IndexOf(e interface{}) int {
	for i, en := 0, ll.head; i < ll.size && en != nil; i, en = i+1, en.next {
		if en.value == e {
			return i
		}
	}
	return -1
}

func (ll *LinkedList) Contains(e interface{}) bool {
	return ll.IndexOf(e) > -1
}

func (ll *LinkedList) ContainsAll(c Collection) bool {
	it := c.GetIterator()
	findAll := true
	for it.HasNext() {
		if ll.IndexOf(it.Next()) < 0 {
			findAll = false
		}
	}
	return findAll
}

func (ll *LinkedList) Get(p int) interface{} {
	ll.checkIndex(p)
	en := ll.head
	for j := 0; en != nil && j < p; j, en = j+1, en.next {
	}
	return en.value
}

func (ll *LinkedList) Empty() bool {
	return ll.size == 0
}

func (ll *LinkedList) Remove(e interface{}) bool {
	en := ll.head
	var f *entry
	for ; en != nil; en = en.next {
		if en.value == e {
			f = en
			break
		}
	}
	return ll.unlink(f)
}

func (ll *LinkedList) RemoveAt(p int) {
	ll.checkIndex(p)
	en := ll.head
	for j := 0; en != nil && j < p; j, en = j+1, en.next {
	}
	ll.unlink(en)
}

func (ll *LinkedList) RemoveAll(c Collection) bool {
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

func (ll *LinkedList) Set(p int, e interface{}) {
	ll.checkIndex(p)
	en := ll.head
	for j := 0; en != nil && j < p; j, en = j+1, en.next {
	}
	en.value = e
	ll.modCount++
}

func (ll *LinkedList) RetainAll(c Collection) bool {
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

func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) ToSlice() []interface{} {
	var dest []interface{}
	entry := ll.head
	for entry != nil {
		dest = append(dest, entry.value)
		entry = entry.next
	}
	return dest
}

func (ll *LinkedList) GetIterator() Iterator {
	return &lIterator{
		next:             ll.head,
		expectedModCount: ll.modCount,
		list:             ll,
		index:            -1,
		lastReturn:       nil,
	}
}

type lIterator struct {
	index            int
	lastReturn       *entry
	next             *entry
	expectedModCount int
	list             *LinkedList
}

func (li *lIterator) HasNext() bool {
	return li.index < li.list.size-1
}

func (li *lIterator) Next() interface{} {
	li.checkModStatus()
	if !li.HasNext() {
		panic("no such elem")
	}
	li.lastReturn = li.next
	li.next = li.next.next
	li.index++
	return li.lastReturn.value
}

func (li *lIterator) Remove() {
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

func (li *lIterator) Set(e interface{}) {
	if li.lastReturn == nil {
		li.invalidLastReturn()
	}
	li.checkModStatus()
	li.lastReturn.value = e
}

func (li *lIterator) Add(e interface{}) {
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

func (li *lIterator) checkModStatus() {
	if li.expectedModCount != li.list.modCount {
		panic("list has been modified during Iterate")
	}
}

func (li *lIterator) invalidLastReturn() {
	panic("last returns is nil")
}
