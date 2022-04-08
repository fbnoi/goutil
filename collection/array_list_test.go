package collection

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestArrayList_Add(t *testing.T) {
	list := &ArrayList{}
	list.Add(1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, true, list.Contains(1))
}

func TestArrayList_AddAt(t *testing.T) {
	list := &ArrayList{}
	list.Add(1)
	list.Add(1)
	list.Add(1)
	list.AddAt(1, 2)
	assert.Equal(t, 2, list.Get(1))
	list.AddAt(0, 0)
	assert.Equal(t, 0, list.Get(0))
	list.AddAt(list.Size()-1, 99)
	assert.Equal(t, 99, list.Get(list.Size()-2))
}

func TestArrayList_AddAll(t *testing.T) {
	list := &ArrayList{}
	l := &ArrayList{}
	l.Add(1)
	l.Add(2)
	list.AddAll(l)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, true, list.Contains(1))
	assert.Equal(t, true, list.Contains(2))
}

func TestArrayList_Clear(t *testing.T) {
	list := &ArrayList{}
	list.Add(1)
	list.Add(1)
	assert.Equal(t, 2, list.Size())
	list.Clear()
	assert.Equal(t, 0, list.Size())
}

func TestArrayList_Contains(t *testing.T) {
	list := &ArrayList{}
	list.Add(1)
	assert.Equal(t, true, list.Contains(1))
	assert.Equal(t, false, list.Contains(2))
}

func TestArrayList_ContainsAll(t *testing.T) {
	list := &ArrayList{}
	list.Add(1)
	list.Add(2)
	l := &ArrayList{}
	l.Add(1)
	assert.Equal(t, true, list.ContainsAll(l))
	l.Add(2)
	assert.Equal(t, true, list.ContainsAll(l))
	l.Add(3)
	assert.Equal(t, false, list.ContainsAll(l))
}

func TestArrayList_TsEmpty(t *testing.T) {
	list := &ArrayList{}
	assert.Equal(t, true, list.Empty())
	list.Add(1)
	assert.Equal(t, false, list.Empty())
}

func TestArrayList_Remove(t *testing.T) {
	list := &ArrayList{}
	assert.Equal(t, false, list.Remove(1))
	list.Add(1)
	list.Add(2)
	assert.Equal(t, true, list.Remove(1))
	assert.Equal(t, true, list.Contains(2))
	assert.Equal(t, false, list.Contains(1))
	assert.Equal(t, 1, list.Size())
}

func TestArrayList_RemoveAll(t *testing.T) {
	list := &ArrayList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	l := &ArrayList{}
	l.Add(1)
	l.Add(2)
	assert.Equal(t, true, list.RemoveAll(l))
	l.Add(3)
	assert.Equal(t, true, list.RemoveAll(l))
	assert.Equal(t, false, list.RemoveAll(l))
	assert.Equal(t, true, list.Empty())
}

func TestArrayList_RetainAll(t *testing.T) {
	list := &ArrayList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	l := &ArrayList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	assert.Equal(t, false, list.RetainAll(l))
	l.Remove(1)
	assert.Equal(t, true, list.RetainAll(l))
	assert.Equal(t, false, list.Contains(1))
	assert.Equal(t, true, list.ContainsAll(l))
}

func TestArrayList_Size(t *testing.T) {
	list := &ArrayList{}
	assert.Equal(t, 0, list.Size())
	list.Add(1)
	assert.Equal(t, 1, list.Size())
}

func TestArrayList_ToSlice(t *testing.T) {
	list := &ArrayList{}
	list.Add(1)
	assert.Equal(t, 1, len(list.ToSlice()))
}

func TestAlIterator_HasNext(t *testing.T) {
	list := &ArrayList{}
	assert.Equal(t, false, list.GetIterator().HasNext())
	list.Add(1)
	it := list.GetIterator()
	assert.Equal(t, true, it.HasNext())
	it.Next()
	assert.Equal(t, false, it.HasNext())
}

func TestAlIterator_Next(t *testing.T) {
	list := &ArrayList{}
	list.Add(1)
	list.Add(2)
	it := list.GetIterator()
	assert.Equal(t, 1, it.Next())
	assert.Equal(t, 2, it.Next())
	assert.PanicMatches(t, func() { it.Next() }, "list bounds out of range")
}

func TestAlIterator_Remove(t *testing.T) {
	list := &ArrayList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	it := list.GetIterator()
	assert.Equal(t, 1, it.Next())
	it.Remove()
	assert.Equal(t, false, list.Contains(1))
	assert.Equal(t, 2, it.Next())
	assert.Equal(t, 3, it.Next())
	assert.Equal(t, false, it.HasNext())
	list.Add(1)
	assert.PanicMatches(t, func() { it.HasNext() }, "list has been modified during Iterate")
}
