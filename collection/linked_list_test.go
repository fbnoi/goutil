package collection

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLinkedList_Add(t *testing.T) {
	ll := &LinkedList[int]{}
	assert.Equal(t, 0, ll.Size())
	ll.Add(1)
	assert.Equal(t, 1, ll.Size())
	assert.Equal(t, 1, ll.modCount)
	ll.Add(2)
	assert.Equal(t, 2, ll.Size())
	assert.Equal(t, 2, ll.modCount)
}

func TestLinkedList_AddAt(t *testing.T) {
	ll := &LinkedList[int]{}
	assert.Equal(t, 0, ll.Size())
	ll.Add(1)
	ll.Add(3)
	assert.Equal(t, 2, ll.Size())
	assert.Equal(t, 2, ll.modCount)
	ll.AddAt(0, 0)
	assert.Equal(t, 3, ll.Size())
	assert.Equal(t, 3, ll.modCount)
	assert.Equal(t, 1, ll.Get(1))
	assert.Equal(t, 3, ll.Get(2))
}

func TestLinkedList_AddAll(t *testing.T) {
	ll := &LinkedList[int]{}
	ls := &LinkedList[int]{}
	ls.Add(1)
	ls.Add(2)
	ll.AddAll(ls)
	assert.Equal(t, 2, ll.Size())
	assert.Equal(t, 1, ll.Get(0))
	assert.Equal(t, 2, ll.Get(1))
	assert.Equal(t, 2, ll.modCount)
}

func TestLinkedList_Clear(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Add(1)
	assert.Equal(t, 1, ll.Size())
	assert.Equal(t, 1, ll.modCount)
	ll.Clear()
	assert.Equal(t, 0, ll.Size())
	assert.Equal(t, 2, ll.modCount)
}

func TestLinkedList_Contains(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Add(1)
	assert.Equal(t, true, ll.Contains(1))
}

func TestLinkedList_ContainsAll(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ls := &LinkedList[int]{}
	ls.Add(2)
	assert.Equal(t, true, ll.ContainsAll(ls))
	ls.Add(4)
	assert.Equal(t, false, ll.ContainsAll(ls))
}

func TestLinkedList_Get(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	assert.Equal(t, 1, ll.Get(0))
	assert.Equal(t, 2, ll.Get(1))
	assert.Equal(t, 3, ll.Get(2))
	assert.PanicMatches(t, func() {
		ll.Get(3)
	}, "list bounds out of range")
}

func TestLinkedList_Empty(t *testing.T) {
	ll := &LinkedList[int]{}
	assert.Equal(t, true, ll.Empty())
	ll.Add(1)
	assert.Equal(t, false, ll.Empty())
}

func TestLinkedList_Remove(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Add(1)
	assert.Equal(t, false, ll.Remove(2))
	assert.Equal(t, true, ll.Remove(1))
	assert.Equal(t, true, ll.Empty())
	assert.Equal(t, false, ll.Contains(1))
}

func TestLinkedList_RemoveAt(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.RemoveAt(0)
	assert.Equal(t, false, ll.Contains(1))
	assert.Equal(t, 2, ll.Size())
	assert.Equal(t, 2, ll.Get(0))
	assert.Equal(t, 3, ll.Get(1))
	assert.PanicMatches(t, func() {
		ll.RemoveAt(3)
	}, "list bounds out of range")
}

func TestLinkedList_RemoveAll(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Add(1)
	ll.Add(1)
	ll.Add(2)
	ls := &ArrayList[int]{}
	ls.Add(1)
	ls.Add(2)
	assert.Equal(t, true, ll.RemoveAll(ls))
	assert.Equal(t, 0, ll.Size())
}

func TestLinkedList_Set(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	assert.Equal(t, 2, ll.Get(1))
	ll.Set(1, 3)
	assert.Equal(t, 3, ll.Get(1))
}

func TestLinkedList_RetainAll(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ls := &LinkedList[int]{}
	ls.Add(1)
	ls.Add(2)
	assert.Equal(t, 4, ll.Size())
	ll.RetainAll(ls)
	assert.Equal(t, 2, ll.Size())
	assert.Equal(t, true, ll.ContainsAll(ls))
}

func TestLinkedList_ITHasNext(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Add(1)
	it := ll.GetIterator()
	assert.Equal(t, true, it.HasNext())
	assert.Equal(t, 1, it.Next())
	assert.Equal(t, false, it.HasNext())
}

func TestLinkedList_ITNext(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	it := ll.GetIterator()
	assert.Equal(t, 1, it.Next())
	assert.Equal(t, 2, it.Next())
	assert.Equal(t, 3, it.Next())
	assert.PanicMatches(t, func() {
		it.Next()
	}, "no such elem")
}

func TestLinkedList_ITRemove(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	it := ll.GetIterator()
	assert.Equal(t, 1, it.Next())
}
